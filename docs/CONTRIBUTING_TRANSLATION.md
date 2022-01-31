# Translation Contribution Guide

So, you're willing to help the Grule-Rule-Engine opensource community to translate our documentation? 
We deeply honoured and greatly appreciate your generosity of your time for this project. 

Please read the following guideline when you're translating the documentations.

## Working with the project

We are following common github contribution practice. 

1. Fork the project into your own github repository.
2. Clone the project from your repostory into your local computer.
3. From your local copy, make an `upstream` remote repository to the original `hyperjump/grule-rule-engine`.
4. From your local copy, create a new branch for your translation project.
5. Make a regular `pull` from the upstream (incase other people work on the same language as you do).
6. Build-up your translation and commit-push to your own repository until a single file is done.
7. Always check your translation result from your repository github (to see how it looks like)
8. Make a PR from your repository/branch.

If the above list still unclear for you, you can always visit [First Contributor manual](https://github.com/firstcontributions/first-contributions) 

## Understanding the documentation layout

Documentations were organized by each of their respective ISO language code.

The directory structure of translations is like the following

```text
grule-rule-engine
  |
  +-- docs
       |
       +-- [iso language code]
            |
            +-- [documentation .md files]
```

Currently it may look something like...

```text
grule-rule-engine
  |
  +-- docs
       |
       +-- cn
       |    +-- About_cn.md
       |    +-- Bencmarking_cn.md
       |    +-- (..more..)_cn.md
       |
       +-- de
       |    +-- About_de.md
       |    +-- Bencmarking_de.md
       |    +-- (..more..)_de.md
       |
       +-- en
       |    +-- About_en.md
       |    +-- Bencmarking_en.md
       |    +-- (..more..)_en.md
       |
       +-- (..more..)
```

You may find that your target language `it` is already there, means that some translation effort already
started for `Italian`. You can still contribute to the document by fixing miss typed word or incorrect grammars.

## Creating your new language translation directory

If you don't see your language ISO code under the `docs` folder, you can create a new directory with the ISO Code as it's name.
Then you should copy the content of the `en` directory to your directory.

For the sake of this documentation, supposed you want to translate into Italian (ISO code is `it`)

Copy the directory (and its content) `grule-rule-engine/docs/en` into `grule-rule-engine/docs/it`.

Why copy from `en (english)` ? Well, english is where code contributors of grule makes their documentation. 
So you can always expect that the most updated documentation is always in english.

After you copy the `en` folder, you should rename the file within your language folder to code with your language ISO.
For example (`About_en.md` to `About_it.md`, `Benchmarking_en.md` to `Benchmarking_it.md` and so on)
Now you're ready to make your translation.

```text
grule-rule-engine
  |
  +-- docs
       |
       +-- cn
       |    +-- About_cn.md
       |    +-- Bencmarking_cn.md
       |    +-- (..more..)_cn.md
       |
       +-- de
       |    +-- About_de.md
       |    +-- Bencmarking_de.md
       |    +-- (..more..)_de.md
       |
       +-- en
       |    +-- About_en.md
       |    +-- Bencmarking_en.md
       |    +-- (..more..)_en.md
       |
       +-- (..more..)
       |
       +-- it
            +-- About_it.md
            +-- Bencmarking_it.md
            +-- (..more..)_it.md
```

## Translating to your language

1. Translate any readable text component. Readable means that any text you see from github page. Including.
   1. The menu titles
   2. "Gopher Holds The Rules" wordings
   3. All the titles and paragraphs as long as deliver the same original meaning. Use the character set applicable for your language.
   4. You might add some more explanation in addition to the current original documentation to make it easier for reader to understand the document context.
2. Do not translate :
   1. Code snippets, those starts and ends with 3 *backquotes*
   2. Names ("Grule", people names) into any other form.
   3. Known abbreviations (e.g. "GRL", "DRL", "ISO", "DSL" etc)

## Links Integrity

As translator AND contributor, you must maintain the integrity of links in your translated pages.

- Menu links between pages in the same language.
- Menu links between pages in to other languages.
- Menu links from other language pages to your language.

This link integrity might get very complex as more language opted in, but we all must courageously fix 
missing and wrong links, even its not in your language.

**PR will never gets approved when we humanely see a broken or wrong links.**

### Tell other that a file is being translated.

Let your readers and contributors know that a file still being translated. Thus for all `.md` file that is not fully translated, please include the following text above the menu...

```markdown
---

:construction:
__THIS PAGE IS BEING TRANSLATED__
:construction:

:construction_worker: Contributors are invited. Please read [CONTRIBUTING](../../CONTRIBUTING.md) and [CONTRIBUTING TRANSLATION](../CONTRIBUTING_TRANSLATION.md) guidelines. 

:vulcan_salute: Please remove this note once you're done translating.

---
```

When it's done, you can remove them.

## Creating PR

1. You can create a PR even for partially finished page. As long as the "Construction" notice still there.
2. Check your PR status often, they might get conflicted overtime thus need you're attention to fix.
3. YOU MUST prefix your PR title with `[LANG/<iso code>]` text, and the title should be in english. E.g.
   1. `[LANG/en] Fixed wrong grammar`
   2. `[LANG/cn] Refining explanation that talk execution cycle`

# Finally

Thank you very much for your contribution and HAPPY TRANSLATING !!!