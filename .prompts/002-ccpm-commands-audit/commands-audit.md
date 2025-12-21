# CCPM Commands Audit Results

```xml
<audit_results type="commands">
  <summary>
    <total_components>52</total_components>
    <critical_issues>48</critical_issues>
    <warning_issues>11</warning_issues>
    <info_issues>2</info_issues>
    <pass_count>23</pass_count>
    <fail_count>29</fail_count>
  </summary>

  <batch number="1" name="root">
    <component name="code-rabbit">
      <path>ccpm/commands/code-rabbit.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>code-rabbit.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Process CodeRabbit review comments with context-aware discretion"</fix>
        </issue>
        <issue severity="critical">
          <location>code-rabbit.md:5-120</location>
          <description>Missing all 3 required XML tags (objective, process, success_criteria). Uses pure markdown structure with # headings</description>
          <fix>Wrap content in XML tags: objective, process, success_criteria. Remove markdown # headings</fix>
        </issue>
        <issue severity="warning">
          <location>code-rabbit.md:5-120</location>
          <description>Uses markdown # and ## headings instead of XML structure throughout file</description>
          <fix>Convert markdown structure to XML tags</fix>
        </issue>
      </issues>
    </component>

    <component name="context">
      <path>ccpm/commands/context.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>context.md:6-7</location>
          <description>Missing all 3 required XML tags (objective, process, success_criteria). Body is just one line: "Invoke the ccpm-context skill for: $ARGUMENTS"</description>
          <fix>Add proper XML structure: objective explaining purpose, process with steps, success_criteria</fix>
        </issue>
      </issues>
    </component>

    <component name="doctor">
      <path>ccpm/commands/doctor.md</path>
      <status>pass</status>
      <issues>
        <issue severity="warning">
          <location>doctor.md:10-51</location>
          <description>Uses ## and ### markdown headings inside process XML tag (lines 11, 13, 22, 32, 40, 49)</description>
          <fix>Remove markdown headings from inside XML; use numbered lists or plain text with bold for structure</fix>
        </issue>
      </issues>
    </component>

    <component name="prompt">
      <path>ccpm/commands/prompt.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>prompt.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Execute prompt from ephemeral command file for complex inputs"</fix>
        </issue>
        <issue severity="critical">
          <location>prompt.md:5-9</location>
          <description>Missing all 3 required XML tags. Uses markdown # heading instead</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
      </issues>
    </component>

    <component name="re-init">
      <path>ccpm/commands/re-init.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>re-init.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Enhance CLAUDE.md file with rules from .claude/CLAUDE.md"</fix>
        </issue>
        <issue severity="critical">
          <location>re-init.md:5-9</location>
          <description>Missing all 3 required XML tags. Uses markdown # heading instead</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
      </issues>
    </component>

    <component name="self-update">
      <path>ccpm/commands/self-update.md</path>
      <status>pass</status>
      <issues>
        <issue severity="warning">
          <location>self-update.md:10-42</location>
          <description>Uses ## and ### markdown headings inside process XML tag</description>
          <fix>Remove markdown headings from inside XML; use numbered lists or plain text</fix>
        </issue>
      </issues>
    </component>

    <component name="setup">
      <path>ccpm/commands/setup.md</path>
      <status>pass</status>
      <issues>
        <issue severity="warning">
          <location>setup.md:10-56</location>
          <description>Uses ## and ### markdown headings inside process XML tag</description>
          <fix>Remove markdown headings from inside XML; use numbered lists or plain text</fix>
        </issue>
      </issues>
    </component>

    <component name="uninstall">
      <path>ccpm/commands/uninstall.md</path>
      <status>pass</status>
      <issues>
        <issue severity="warning">
          <location>uninstall.md:10-57</location>
          <description>Uses ## and ### markdown headings inside process XML tag</description>
          <fix>Remove markdown headings from inside XML; use numbered lists or plain text</fix>
        </issue>
      </issues>
    </component>

    <component name="version">
      <path>ccpm/commands/version.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure without markdown headings -->
      </issues>
    </component>
  </batch>

  <batch number="2" name="context-testing">
    <component name="context/create">
      <path>ccpm/commands/context/create.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="context/prime">
      <path>ccpm/commands/context/prime.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="context/update">
      <path>ccpm/commands/context/update.md</path>
      <status>pass</status>
      <issues>
        <issue severity="info">
          <location>context/update.md:28</location>
          <description>$ARGUMENTS appears alone at end of file, outside any XML tag</description>
          <fix>Move $ARGUMENTS into the context or process tag where it's referenced</fix>
        </issue>
      </issues>
    </component>

    <component name="testing/prime">
      <path>ccpm/commands/testing/prime.md</path>
      <status>pass</status>
      <issues>
        <issue severity="warning">
          <location>testing/prime.md:1-4</location>
          <description>Uses $ARGUMENTS at line 27 but no argument-hint in YAML frontmatter</description>
          <fix>Add argument-hint: "[options]" if arguments are expected</fix>
        </issue>
        <issue severity="info">
          <location>testing/prime.md:27</location>
          <description>$ARGUMENTS appears alone at end of file, outside any XML tag</description>
          <fix>Move $ARGUMENTS into context tag or remove if not needed</fix>
        </issue>
      </issues>
    </component>

    <component name="testing/run">
      <path>ccpm/commands/testing/run.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure with $ARGUMENTS in context tag -->
      </issues>
    </component>
  </batch>

  <batch number="3" name="epic">
    <component name="epic-close">
      <path>ccpm/commands/pm/epic-close.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-decompose">
      <path>ccpm/commands/pm/epic-decompose.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-edit">
      <path>ccpm/commands/pm/epic-edit.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-list">
      <path>ccpm/commands/pm/epic-list.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>epic-list.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="epic-merge">
      <path>ccpm/commands/pm/epic-merge.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-oneshot">
      <path>ccpm/commands/pm/epic-oneshot.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-refresh">
      <path>ccpm/commands/pm/epic-refresh.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-show">
      <path>ccpm/commands/pm/epic-show.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>epic-show.md:5-7</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="epic-start">
      <path>ccpm/commands/pm/epic-start.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="epic-start-worktree">
      <path>ccpm/commands/pm/epic-start-worktree.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>epic-start-worktree.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Launch parallel agents to work on epic tasks in a shared worktree"</fix>
        </issue>
        <issue severity="critical">
          <location>epic-start-worktree.md:5-222</location>
          <description>Missing all 3 required XML tags. Uses pure markdown structure with # headings</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
        <issue severity="warning">
          <location>epic-start-worktree.md:5-222</location>
          <description>Uses $ARGUMENTS throughout but no argument-hint in YAML frontmatter</description>
          <fix>Add argument-hint: &lt;epic_name&gt;</fix>
        </issue>
      </issues>
    </component>

    <component name="epic-status">
      <path>ccpm/commands/pm/epic-status.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>epic-status.md:5-7</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="epic-sync">
      <path>ccpm/commands/pm/epic-sync.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>
  </batch>

  <batch number="4" name="issue">
    <component name="issue-analyze">
      <path>ccpm/commands/pm/issue-analyze.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="issue-close">
      <path>ccpm/commands/pm/issue-close.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="issue-edit">
      <path>ccpm/commands/pm/issue-edit.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="issue-reopen">
      <path>ccpm/commands/pm/issue-reopen.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="issue-show">
      <path>ccpm/commands/pm/issue-show.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>issue-show.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Display issue and sub-issues with detailed information"</fix>
        </issue>
        <issue severity="critical">
          <location>issue-show.md:5-91</location>
          <description>Missing all 3 required XML tags. Uses pure markdown structure with # headings</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
        <issue severity="warning">
          <location>issue-show.md:5-91</location>
          <description>Uses $ARGUMENTS throughout but no argument-hint in YAML frontmatter</description>
          <fix>Add argument-hint: &lt;issue_number&gt;</fix>
        </issue>
      </issues>
    </component>

    <component name="issue-start">
      <path>ccpm/commands/pm/issue-start.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="issue-status">
      <path>ccpm/commands/pm/issue-status.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="issue-sync">
      <path>ccpm/commands/pm/issue-sync.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>
  </batch>

  <batch number="5" name="prd">
    <component name="prd-edit">
      <path>ccpm/commands/pm/prd-edit.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="prd-list">
      <path>ccpm/commands/pm/prd-list.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>prd-list.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="prd-new">
      <path>ccpm/commands/pm/prd-new.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="prd-parse">
      <path>ccpm/commands/pm/prd-parse.md</path>
      <status>pass</status>
      <issues>
        <!-- No issues - proper XML structure -->
      </issues>
    </component>

    <component name="prd-status">
      <path>ccpm/commands/pm/prd-status.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>prd-status.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>
  </batch>

  <batch number="6" name="general-pm">
    <component name="blocked">
      <path>ccpm/commands/pm/blocked.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>blocked.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="clean">
      <path>ccpm/commands/pm/clean.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>clean.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Clean up completed work and archive old epics"</fix>
        </issue>
        <issue severity="critical">
          <location>clean.md:5-102</location>
          <description>Missing all 3 required XML tags. Uses pure markdown structure with # headings</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
      </issues>
    </component>

    <component name="help">
      <path>ccpm/commands/pm/help.md</path>
      <status>fail</status>
      <issues>
        <issue severity="warning">
          <location>help.md:6-14</location>
          <description>Has objective and success_criteria but missing required process tag. Bash command is between them, not in a tag</description>
          <fix>Add process tag containing the bash command invocation</fix>
        </issue>
      </issues>
    </component>

    <component name="import">
      <path>ccpm/commands/pm/import.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>import.md:7-100</location>
          <description>Missing all 3 required XML tags. Uses pure markdown structure with # headings</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
      </issues>
    </component>

    <component name="in-progress">
      <path>ccpm/commands/pm/in-progress.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>in-progress.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="init">
      <path>ccpm/commands/pm/init.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>init.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="next">
      <path>ccpm/commands/pm/next.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>next.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="search">
      <path>ccpm/commands/pm/search.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>search.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
        <issue severity="warning">
          <location>search.md:6</location>
          <description>Uses $ARGUMENTS in bash command but no argument-hint in YAML frontmatter</description>
          <fix>Add argument-hint: &lt;search_query&gt;</fix>
        </issue>
      </issues>
    </component>

    <component name="standup">
      <path>ccpm/commands/pm/standup.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>standup.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="status">
      <path>ccpm/commands/pm/status.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>status.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>

    <component name="sync">
      <path>ccpm/commands/pm/sync.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>sync.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Full bidirectional sync between local and GitHub"</fix>
        </issue>
        <issue severity="critical">
          <location>sync.md:5-82</location>
          <description>Missing all 3 required XML tags. Uses pure markdown structure with # headings</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
        <issue severity="warning">
          <location>sync.md:11-14</location>
          <description>Accepts optional epic_name argument but no argument-hint in YAML frontmatter</description>
          <fix>Add argument-hint: "[epic_name]"</fix>
        </issue>
      </issues>
    </component>

    <component name="test-reference-update">
      <path>ccpm/commands/pm/test-reference-update.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>test-reference-update.md:1-3</location>
          <description>Missing required 'description' field in YAML frontmatter</description>
          <fix>Add description: "Test the task reference update logic used in epic-sync"</fix>
        </issue>
        <issue severity="critical">
          <location>test-reference-update.md:5-134</location>
          <description>Missing all 3 required XML tags. Uses pure markdown structure with # headings</description>
          <fix>Wrap content in objective, process, success_criteria XML tags</fix>
        </issue>
      </issues>
    </component>

    <component name="validate">
      <path>ccpm/commands/pm/validate.md</path>
      <status>fail</status>
      <issues>
        <issue severity="critical">
          <location>validate.md:5-6</location>
          <description>Missing all 3 required XML tags. Body is just bash script invocation</description>
          <fix>Add objective, process, success_criteria XML tags wrapping the bash command</fix>
        </issue>
      </issues>
    </component>
  </batch>

  <pattern_issues>
    <pattern name="bash-delegator-no-xml" count="14">
      <description>Commands that only invoke a bash script without any XML structure</description>
      <affected_commands>epic-list, epic-show, epic-status, prd-list, prd-status, blocked, in-progress, init, next, search, standup, status, validate</affected_commands>
      <fix>Add minimal XML wrapper: objective (what it does), process (runs script), success_criteria (what success looks like)</fix>
    </pattern>

    <pattern name="markdown-headings-in-xml" count="5">
      <description>Commands with proper XML tags but markdown ## headings inside them</description>
      <affected_commands>doctor, self-update, setup, uninstall</affected_commands>
      <fix>Remove markdown headings from inside XML tags; use numbered lists or bold text for structure</fix>
    </pattern>

    <pattern name="pure-markdown-structure" count="8">
      <description>Commands with extensive markdown content but no XML structure at all</description>
      <affected_commands>code-rabbit, epic-start-worktree, issue-show, clean, import, sync, test-reference-update</affected_commands>
      <fix>Restructure entire content into XML format with objective, process, success_criteria tags</fix>
    </pattern>

    <pattern name="missing-argument-hint" count="5">
      <description>Commands using $ARGUMENTS but lacking argument-hint in frontmatter</description>
      <affected_commands>epic-start-worktree, issue-show, search, sync, testing/prime</affected_commands>
      <fix>Add argument-hint field to YAML frontmatter describing expected arguments</fix>
    </pattern>

    <pattern name="missing-description" count="7">
      <description>Commands lacking required description field in YAML frontmatter</description>
      <affected_commands>code-rabbit, prompt, re-init, epic-start-worktree, issue-show, clean, sync, test-reference-update</affected_commands>
      <fix>Add description field with clear, concise explanation of what command does</fix>
    </pattern>
  </pattern_issues>

  <metadata>
    <audit_date>2025-12-21</audit_date>
    <auditor>manual-audit (6 sequential batches)</auditor>
    <audit_criteria>
      <criterion name="yaml-compliance">description required, argument-hint if args, allowed-tools if restrictions</criterion>
      <criterion name="xml-structure">required tags: objective, process, success_criteria</criterion>
      <criterion name="no-markdown-in-xml">no ## or ### headings inside XML tags</criterion>
      <criterion name="argument-handling">$ARGUMENTS requires argument-hint in frontmatter</criterion>
      <criterion name="dynamic-context">proper ! backtick and @ file syntax</criterion>
    </audit_criteria>
  </metadata>
</audit_results>
```
