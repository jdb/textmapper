#include "builder.h"
#include "selector.h"

#include <iostream>
#include <string>

void DisplayNode(const json::Node* node, int depth) {
  if (!node) {
    return;
  }
  
  // Safety check for infinite recursion
  if (depth > 100) {
    std::cerr << "ERROR: Maximum depth exceeded! Possible cycle detected." << std::endl;
    return;
  }

  // Print indentation
  for (int i = 0; i < depth; i++) {
    std::cout << "  ";
  }

  // Print node type and text range
  std::cout << node->Type() << " [" << node->Offset() << "-" << node->Endoffset() << "]";
  
  // Print node text if it's a leaf or short
  std::string text = node->Text();
  if (text.size() < 50 && !text.empty()) {
    std::cout << " \"" << text << "\"";
  }
  std::cout << std::endl;

  std::vector<const json::Node*> children = node->Children(json::Any);
  
  // Safety check for excessive children (indicates cycle)
  if (children.size() > 1000) {
    std::cerr << "ERROR: Excessive children count indicates possible cycle!" << std::endl;
    return;
  }
  
  // Recursively display children
  for (const json::Node* child : children) {
    if (child) {
      DisplayNode(child, depth + 1);
    }
  }
}

int main(int argc, char* argv[]) {
  if (argc < 2) {
    std::cerr << "Usage: " << argv[0] << " <json_string>" << std::endl;
    return 1;
  }

  std::string input = argv[1];
  json::Tree tree = json::Parse(input);

  const json::Node* root = tree.Root();
  if (!root) {
    std::cerr << "Failed to parse or empty tree" << std::endl;
    return 1;
  }

  DisplayNode(root, 0);
  return 0;
}
