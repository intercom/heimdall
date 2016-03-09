# Heimdall
Heimdall is a tool for finding specific text patterns that you don't want in your code. It requires [The Silver Searcher](https://github.com/ggreer/the_silver_searcher) to be installed.

# Rules
Rules will be read from a `heimdall.json` file in your working directory. The file uses the following format:
```json
{
  "rules":[
    {
      "name": "Rule Name",
      "pattern": ".+The pattern you want to search for.*",
      "description": "An explanation of what the rule is and why it's used",
      "flags":"--ignore *.unwanted-extension",
      "path":"repo_dir"
    }
  ]
}
```
