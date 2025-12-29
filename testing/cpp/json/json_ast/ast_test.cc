#include "builder.h"
#include "selector.h"

#include <cstdint>
#include <functional>
#include <ostream>
#include <tuple>
#include <unordered_set>
#include <vector>

#include "absl/strings/string_view.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"
#include "cpp/markup/markup.h"

namespace json {
namespace {

struct AstTestCase {
  NodeType nt;
  std::vector<absl::string_view> inputs;
};
inline std::ostream& operator<<(std::ostream& os, const AstTestCase& t) {
  return os << "{" << t.nt << ", " << t.inputs.size() << " cases}";
}

TEST(AstTest, ParseSimpleValue) {
  Tree tree = Parse("1.0");
  const Node* root = tree.Root();
  
  ASSERT_NE(root, nullptr);
  EXPECT_EQ(root->Type(), NodeType::JSONText);
  
  const Node* json_value = root->Child(JSONValue);
  ASSERT_NE(json_value, nullptr);
  EXPECT_EQ(json_value->Type(), NodeType::JSONValue);
}

TEST(AstTest, NodeTypes) {
  const AstTestCase tests[] = {
      {NodeType::JSONEmptyObject,
       {
           R"(«{}»)",
           R"(«{ /* comment */ }»)",
           R"({"aa": «{}» })",
       }},
      {NodeType::JSONObject,
       {
           R"(«{ "a" : "b" }»)",
           R"(«{ "a" : ["b"] }»)",
           R"(«{ "a" : {} }»)",
           R"(«{ "a" : «{"q":B}» }»)",
       }},
      {NodeType::JSONArray,
       {
           R"({ "a" : «["b"]» })",
           R"( «[]» )",
       }},
      {NodeType::JSONText,
       {
           R"(«{ "a" : ["b", A] }»)",
           R"( «"aa"» )",
           R"( «A» )",
       }},
      {NodeType::JSONMember,
       {
           R"([{ «"a" : ["b"]», «"q":[]» }])",
       }},
      {NodeType::JSONValue,
       {
           R"(«{ "a" : «[«"b"»]» }»)",
           R"( «"aa"» )",
       }},
      {NodeType::MultiLineComment,
       {
           R"({ "a"«/* abc */» : [] })",
       }},
      {NodeType::JSONString,
       {
           R"({ «"a"» : [«"b"»] })",
       }},
      {NodeType::SyntaxProblem,
       {
           R"({ «"a" "b"» })",
       }},
  };
  std::unordered_set<NodeType> seen;
  for (auto& test_case : tests) {
    LOG(INFO) << "Running AST tests for " << test_case;
    for (const auto& input : test_case.inputs) {
      auto [want, text] = markup::Parse(input);
      Tree tree = Parse(std::string(text));
      std::vector<markup::Range> got;
      if (tree.Root()) {
        tree.Root()->Traverse([&](const Node* n) -> bool {
          if (n->Type() == test_case.nt) {
            seen.insert(n->Type());
            got.push_back(markup::Range{n->Offset(), n->Endoffset()});
          }
          return true;
        });
      }
      EXPECT_THAT(got, testing::UnorderedElementsAreArray(want))
          << "Node type test failed for " << test_case.nt << " on input:\n"
          << input;
    }
  }
}

}
}
