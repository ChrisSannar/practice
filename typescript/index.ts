import markdownIt from 'markdown-it-ts'

const test = `# Hello World

I am some random text

## I am a second heading
 - I am list item 1
 - I am list item 2
 - I am list item 3

\`\`\`
I am a code block
\`\`\`
`

const md = markdownIt()
const tokens = md.parse(test)

console.log(test)
console.log()
for (const t of tokens) {
  console.log(t.type, t.content)
}
