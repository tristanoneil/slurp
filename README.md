# Slurp 🍜

Slurp is a CLI tool to help me (and maybe you) build a grocery list. It currently supports:

- Provide grocery items line by line or comma separated
- Type done + <return> when you're done
- Your items are sent to an LLM (currently just OpenAI) and are returned based on a configured prompt

### Example Config

This TOML file should be located at `~/.config/slurp/config`

```TOML
prompt = """
You are a helpful assistant specializing in grocery categorization. Prefix each
item in a grocery list with a single emoji representing its store location.
Remember the emoji should represent the category NOT the item. Sort items by
their prefixed emoji. Do not include category headers. If an item's category is
unclear, default to 🛒 pantry. Correct spelling errors and capitalize item
names. The valid categories are: produce, dairy, pantry, frozen, meat,
toiletries. For pantry use 🛒, for produce use 🍑, for meat use 🥩, for frozen
use 🧊, for toiletries use 🧴
"""
```

### Why?

- I wasn't happy with my current workflow for creating grocery lists
- I wanted to learn Go
- I love CLI tools and felt like building one
- I was bored during a long winter break ☃️

### Future

- Send final list to a todo platform (TBD which one)
- Allow for removing items
- Allow for listing items
- Other cool stuff
- Add tests
- Setup CI (for tests and building releases)
