#include "json_lexer.h"

#include <sstream>
#include <vector>

#include "gmock/gmock.h"
#include "gtest/gtest.h"
#include "cpp/markup/markup.h"

namespace json {
namespace {

struct Test {
  std::string name;
  Token want_tok;
  std::vector<std::string> cases;
};

inline std::ostream &operator<<(std::ostream &os, Test t) {
  return os << "{" << t.want_tok << ", " << t.cases.size() << " cases}";
}

const std::vector<Test> tests = {
    {"id", Token::ID, {R"( 0«foo» «barB1»)",}},
    {"string", Token::JSONSTRING, {R"(«"foo"» «"b\nar"» «"α"»)",}},
    {"number", Token::JSONNUMBER, {"«1» «534»", "«1e9» «1.2» «1e-2»",}},
    {"true", Token::TRUE, {"«true»", "/* true */ «true» "}},
    {"false", Token::FALSE, {"  «false» "}},
    {"null", Token::KW_NULL, {"  «null» "}},
    {"lbrace", Token::LBRACE, {"«{»"}},
    {"rbrace", Token::RBRACE, {"«}»"}},
    {"lbrack", Token::LBRACK, {"«[»"}},
    {"rbrack", Token::RBRACK, {"«]»"}},
    {"colon", Token::COLON, {"«:»"}},
    {"comma", Token::COMMA, {"«,»"}},
    {"comment", Token::MULTILINECOMMENT, {"  «/*  asda *** */» bar"}},
    {"char_a", Token::CHAR_A, {"  «A» «α» «A»"}},

    // TODO: handle invalid tokens
    //
    // {"invalid",
    //  Token::INVALID_TOKEN,
    //  {
    //      "«1e» ",
    //      "abc   «/*1e  \n»",
    //      "abc   «\"  »\n",
    //  }},
};

struct RuneOffsetTestCase {
  std::string name;
  std::string json_string;
  int64_t expected_byte_span;
  int64_t expected_char_span;
};

inline std::ostream &operator<<(std::ostream &os, const RuneOffsetTestCase &t) {
  return os << "{" << t.name << ", byte_span=" << t.expected_byte_span
            << ", char_span=" << t.expected_char_span << "}";
}

class LexerTest : public testing::TestWithParam<Test> {};

TEST_P(LexerTest, Token) {
  const auto &param = GetParam();
  for (const auto &input : param.cases) {
    std::vector<markup::Range> want;
    std::string text;
    tie(want, text) = markup::Parse(input);

    Lexer l(text);
    std::vector<markup::Range> got;
    Token tok;
    while ((tok = l.Next()) != Token::EOI) {
      if (tok == param.want_tok) {
        auto loc = l.LastTokenLocation();
        got.push_back(
            markup::Range{loc.begin, loc.end});
      }
    }

    EXPECT_THAT(got, testing::ElementsAreArray(want))
        << "lexer produced " << markup::Create(text, got) << " instead of "
        << markup::Create(text, want);
  }
}

INSTANTIATE_TEST_SUITE_P(Vals, LexerTest, testing::ValuesIn(tests),
                         [](const ::testing::TestParamInfo<Test> &info) {
                           return info.param.name;
                         });

const std::vector<RuneOffsetTestCase> rune_offset_tests = {
    {"2byte_alpha", R"("α")", 4, 3},
    {"3byte_chinese", R"("中")", 5, 3},
    {"4byte_emoji", R"("🎉")", 6, 3},
    {"mixed_chars", R"("aα中🎉")", 12, 6},
};

class RuneOffsetTest : public testing::TestWithParam<RuneOffsetTestCase> {};

TEST_P(RuneOffsetTest, CharLocationWithMultiByteUTF8) {
  const auto &param = GetParam();
  
  Lexer l(param.json_string);
  Token tok = l.Next();
  EXPECT_EQ(tok, Token::JSONSTRING);
  
  auto byte_loc = l.LastTokenLocation();
  Lexer::Location char_loc(l.token_offset_rune_, l.offset_rune_);
  
  int64_t byte_span = byte_loc.end - byte_loc.begin;
  int64_t char_span = char_loc.end - char_loc.begin;
  
  EXPECT_EQ(byte_span, param.expected_byte_span)
      << "Test case: " << param.name;
  EXPECT_EQ(char_span, param.expected_char_span)
      << "Test case: " << param.name;
  EXPECT_GT(byte_span, char_span)
      << "Test case: " << param.name << " - Byte span should be greater than character span due to multi-byte UTF-8 characters";
}

INSTANTIATE_TEST_SUITE_P(RuneOffsets, RuneOffsetTest,
                         testing::ValuesIn(rune_offset_tests),
                         [](const ::testing::TestParamInfo<RuneOffsetTestCase> &info) {
                           return info.param.name;
                         });

struct CountRunesTestCase {
  std::string name;
  std::string input;
  int64_t start_byte;
  int64_t end_byte;
  int64_t expected_rune_count;
  std::string description;
};

inline std::ostream &operator<<(std::ostream &os, const CountRunesTestCase &t) {
  return os << "{" << t.name << ", start=" << t.start_byte << ", end=" << t.end_byte
            << ", expected=" << t.expected_rune_count << "}";
}

const std::vector<CountRunesTestCase> count_runes_tests = {
    // ASCII characters
    {"single_ascii", "a", 0, 1, 1, "Single ASCII character"},
    {"multiple_ascii", "abc", 0, 3, 3, "Multiple ASCII characters"},
    {"partial_ascii_range", "abcdef", 2, 5, 3, "Partial ASCII range"},
    {"empty_range", "abc", 1, 1, 0, "Empty range (start == end)"},
    {"zero_start", "abc", 0, 0, 0, "Zero-length range at start"},
    
    // 2-byte UTF-8 sequences (e.g., Greek alpha: α = 0xCE 0xB1)
    {"single_2byte", "α", 0, 2, 1, "Single 2-byte UTF-8 (Greek alpha)"},
    {"multiple_2byte", "αβ", 0, 4, 2, "Multiple 2-byte sequences"},
    {"partial_2byte_start", "αβ", 0, 1, 1, "Incomplete 2-byte at start (only first byte, treated as invalid)"},
    {"partial_2byte_mid", "αβ", 2, 3, 1, "Incomplete 2-byte in middle (only first byte, treated as invalid)"},
    {"2byte_at_boundary", "aα", 1, 2, 1, "2-byte sequence cut at boundary (only first byte, treated as invalid)"},
    {"2byte_incomplete_end", "α", 0, 1, 1, "2-byte sequence incomplete at end (treated as invalid)"},
    
    // 3-byte UTF-8 sequences (e.g., Chinese: 中 = 0xE4 0xB8 0xAD)
    {"single_3byte", "中", 0, 3, 1, "Single 3-byte UTF-8 (Chinese)"},
    {"multiple_3byte", "中文", 0, 6, 2, "Multiple 3-byte sequences"},
    {"partial_3byte_start", "中", 0, 1, 1, "Incomplete 3-byte at start (only first byte, treated as invalid)"},
    {"partial_3byte_2bytes", "中", 0, 2, 1, "Incomplete 3-byte at 2 bytes (first byte treated as invalid, continuation skipped)"},
    {"3byte_at_boundary", "a中", 1, 2, 1, "3-byte sequence cut at boundary (only first byte, treated as invalid)"},
    {"3byte_incomplete_end", "中", 0, 1, 1, "3-byte sequence incomplete at end (treated as invalid)"},
    {"3byte_incomplete_2bytes", "中", 0, 2, 1, "3-byte sequence incomplete at 2 bytes (first byte treated as invalid, continuation skipped)"},
    
    // 4-byte UTF-8 sequences (e.g., emoji: 🎉 = 0xF0 0x9F 0x8E 0x89)
    {"single_4byte", "🎉", 0, 4, 1, "Single 4-byte UTF-8 (emoji)"},
    {"multiple_4byte", "🎉🎊", 0, 8, 2, "Multiple 4-byte sequences"},
    {"partial_4byte_start", "🎉", 0, 1, 1, "Incomplete 4-byte at start (only first byte, treated as invalid)"},
    {"partial_4byte_2bytes", "🎉", 0, 2, 1, "Incomplete 4-byte at 2 bytes (first byte treated as invalid, continuation skipped)"},
    {"partial_4byte_3bytes", "🎉", 0, 3, 1, "Incomplete 4-byte at 3 bytes (first byte treated as invalid, 2 continuations skipped)"},
    {"4byte_at_boundary", "a🎉", 1, 2, 1, "4-byte sequence cut at boundary (only first byte, treated as invalid)"},
    {"4byte_incomplete_end", "🎉", 0, 1, 1, "4-byte sequence incomplete at end (treated as invalid)"},
    {"4byte_incomplete_2bytes", "🎉", 0, 2, 1, "4-byte sequence incomplete at 2 bytes (first byte treated as invalid, continuation skipped)"},
    {"4byte_incomplete_3bytes", "🎉", 0, 3, 1, "4-byte sequence incomplete at 3 bytes (first byte treated as invalid, 2 continuations skipped)"},
    
    // Continuation bytes (should be skipped)
    {"continuation_byte_only", std::string(1, '\x80'), 0, 1, 0, "Lone continuation byte"},
    {"continuation_bytes", std::string({'\x80', '\x81', '\x82'}), 0, 3, 0, "Multiple continuation bytes"},
    {"continuation_after_valid", std::string({'a', '\x80', '\x81'}), 0, 3, 1, "Continuation bytes after valid char"},
    {"continuation_before_valid", std::string({'\x80', '\x81', 'a'}), 0, 3, 1, "Continuation bytes before valid char"},
    {"continuation_mixed", std::string({'a', '\x80', 'b', '\x81', 'c'}), 0, 5, 3, "Continuation bytes mixed with valid chars"},
    
    // Invalid sequences (treated as single byte)
    {"invalid_high_byte", std::string(1, '\xFF'), 0, 1, 1, "Invalid high byte (0xFF)"},
    {"invalid_sequence", std::string({'\xFE', '\xFF'}), 0, 2, 2, "Invalid UTF-8 sequence"},
    {"invalid_after_valid", std::string({'a', '\xFF'}), 0, 2, 2, "Invalid byte after valid char"},
    
    // Mixed sequences
    {"mixed_ascii_utf8", "aα中🎉", 0, 10, 4, "Mixed ASCII and multi-byte (a=1, α=2, 中=3, 🎉=4 bytes, total=10)"},
    {"mixed_with_continuation", std::string({'a', '\x80', '\xCE', '\xB1'}), 0, 4, 2, "Mixed with continuation bytes"},
    {"mixed_partial_range", "aα中🎉", 1, 9, 3, "Partial range starting mid-sequence"},
    
    // Boundary conditions
    {"beyond_string_end", "abc", 0, 100, 3, "End beyond string size"},
    {"start_beyond_string", "abc", 10, 20, 0, "Start beyond string size"},
    {"start_at_end", "abc", 3, 5, 0, "Start at string end"},
    
    // Edge cases with sequences at boundaries
    {"2byte_split", "αβ", 1, 3, 1, "2-byte sequence split (continuation byte + start of next, continuation skipped, start counted)"},
    {"3byte_split_1", "中", 1, 2, 0, "3-byte sequence split at byte 1-2"},
    {"3byte_split_2", "中", 2, 3, 0, "3-byte sequence split at byte 2-3"},
    {"4byte_split_1", "🎉", 1, 2, 0, "4-byte sequence split at byte 1-2"},
    {"4byte_split_2", "🎉", 2, 3, 0, "4-byte sequence split at byte 2-3"},
    {"4byte_split_3", "🎉", 3, 4, 0, "4-byte sequence split at byte 3-4"},
};

class CountRunesTest : public testing::TestWithParam<CountRunesTestCase> {};

TEST_P(CountRunesTest, CountRunesPaths) {
  const auto &param = GetParam();
  
  int64_t result = countRunes(param.input, param.start_byte, param.end_byte);
  
  EXPECT_EQ(result, param.expected_rune_count)
      << "Test case: " << param.name << " (" << param.description << ")\n"
      << "  Range: [" << param.start_byte << ", " << param.end_byte << ")\n"
      << "  Expected: " << param.expected_rune_count << " runes\n"
      << "  Got: " << result << " runes";
}

INSTANTIATE_TEST_SUITE_P(CountRunesPaths, CountRunesTest,
                         testing::ValuesIn(count_runes_tests),
                         [](const ::testing::TestParamInfo<CountRunesTestCase> &info) {
                           return info.param.name;
                         });

}  // namespace
}  // namespace json

