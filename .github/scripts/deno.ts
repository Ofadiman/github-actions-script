import {parse} from 'https://deno.land/std@0.178.0/flags/mod.ts'

const parsedArgs = parse(Deno.args, {
  string: ['message'],
})
console.log(`parsed message flag: ${parsedArgs.message}`)

const githubOutput = Deno.env.get('GITHUB_OUTPUT') as string
await Deno.writeFile(githubOutput, new TextEncoder().encode(`value=example\n`), {append: true})
