"""
Python implementation of markup parsing, replicating the C++ markup::Parse and markup::Create behavior.

This module provides functions to parse markup format with «text» markers,
extracting expected ranges and cleaning the text, as well as creating
markup strings from text and ranges.
"""


def parse(text):
    """
    Parse markup format «text» and return (ranges, cleaned_text).
    
    The function extracts ranges marked by «text» markers and returns
    both the expected ranges (as tuples of (start, end)) and the cleaned
    text with markers removed.
    
    The ranges are character offsets (not bytes) in the cleaned text.
    Multibyte UTF-8 characters count as a single character.
    
    Args:
        text: Input string with «text» markers
        
    Returns:
        tuple: (list of (start, end) tuples, cleaned_text string)
        
    Raises:
        ValueError: If markup markers are mismatched or invalid
    """
    opening = "«"
    closing = "»"
    out_parts = []  # List of string parts to join
    ranges = []
    stack = []
    
    pos = 0
    i = 0
    text_len = len(text)
    
    while i < text_len:
        # Check for opening marker «
        if text[i:i+len(opening)] == opening:
            # Append text before marker
            out_parts.append(text[pos:i])
            pos = i + len(opening)
            # Push range with start = current output size (in characters), end = -1 (to be filled)
            current_output = ''.join(out_parts)
            stack.append((len(current_output), -1))
            i = pos
            continue
                
        # Check for closing marker »
        if text[i:i+len(closing)] == closing:
            # Append text before marker
            out_parts.append(text[pos:i])
            pos = i + len(closing)
            if not stack:
                raise ValueError(f"unexpected closing guillemets in {text}")
            # Pop range and set end (in characters)
            current_output = ''.join(out_parts)
            start, _ = stack.pop()
            end = len(current_output)
            ranges.append((start, end))
            i = pos
            continue
                
        i += 1
    
    # Append remaining text
    out_parts.append(text[pos:])
    cleaned_text = ''.join(out_parts)
    
    if stack:
        raise ValueError(f"missing closing guillemets in {text}")
        
    return ranges, cleaned_text


def create(text, ranges):
    """
    Create a string with guillemets inserted at the given ranges.
    
    The function inserts « at the start of each range and » at the end.
    
    Args:
        text: Input string
        ranges: List of (start, end) tuples representing character offsets
               (not bytes). Multibyte UTF-8 characters count as a single character.
        
    Returns:
        str: Text with « and » markers inserted at the specified ranges
        
    Raises:
        ValueError: If ranges are invalid (negative, out of bounds, or start > end)
    """
    opening = "«"
    closing = "»"
    
    # Validate ranges
    for start, end in ranges:
        if start < 0 or end < 0:
            raise ValueError(f"range has negative offset: ({start}, {end})")
        if start > end:
            raise ValueError(f"range start > end: ({start}, {end})")
        if end > len(text):
            raise ValueError(f"range end out of bounds: ({start}, {end}) for text of length {len(text)}")
    
    # Create list of brackets with their positions
    brackets = []
    for start, end in ranges:
        brackets.append((start, opening))
        brackets.append((end, closing))
    
    # Sort brackets by offset, with closing brackets (») coming before
    # opening brackets («) at the same position
    def bracket_key(bracket):
        offset, marker = bracket
        # Use 0 for closing, 1 for opening to ensure closing comes first at same offset
        marker_order = 0 if marker == closing else 1
        return (offset, marker_order)
    
    brackets.sort(key=bracket_key)
    
    # Build output string
    out_parts = []
    i = 0
    for offset, marker in brackets:
        if i < offset:
            out_parts.append(text[i:offset])
        out_parts.append(marker)
        i = offset
    
    # Append remaining text
    if i < len(text):
        out_parts.append(text[i:])
    
    return ''.join(out_parts)

