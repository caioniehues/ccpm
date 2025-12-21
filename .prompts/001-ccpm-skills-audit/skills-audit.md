# CCPM Skills Audit Results

```xml
<audit_results type="skills">
  <summary>
    <total_components>6</total_components>
    <critical_issues>26</critical_issues>
    <warning_issues>12</warning_issues>
    <info_issues>4</info_issues>
  </summary>

  <component name="ccpm-context">
    <path>ccpm/skills/ccpm-context/</path>
    <status>fail</status>
    <issues>
      <issue severity="critical">
        <location>SKILL.md:6-49</location>
        <description>Markdown headings inside XML tag. The essential_principles block contains markdown headings (## How Context Management Works, ### Context Structure, etc.) which violates pure XML structure requirements.</description>
        <fix>Replace markdown headings with semantic XML tags like overview, context_structure, frontmatter_requirements, context_lifecycle, key_principles</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md</location>
        <description>Missing required objective tag. Every skill MUST have an objective tag explaining what the skill does and why it matters.</description>
        <fix>Add objective tag at start of body</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md</location>
        <description>Missing required quick_start tag. The intake tag does not replace the need for a quick_start that provides immediate actionable guidance.</description>
        <fix>Add quick_start tag with immediate usage pattern</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="warning">
        <location>workflows/create-context.md:4-110</location>
        <description>Markdown headings inside process tag (## Step 1 through ## Step 6). Workflow files should also use pure XML structure.</description>
        <fix>Replace ## Step N headings with XML tags like step_preflight, step_analysis, step_create, step_validate</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>workflows/prime-context.md:4-129</location>
        <description>Markdown headings inside process tag (## Step 1 through ## Step 8). Same hybrid structure issue.</description>
        <fix>Replace ## Step N headings with semantic XML tags</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>workflows/update-context.md:4-189</location>
        <description>Markdown headings inside process tag (## Step 1 through ## Step 9). Same hybrid structure issue.</description>
        <fix>Replace ## Step N headings with semantic XML tags</fix>
        <pre_identified>true</pre_identified>
      </issue>
    </issues>
  </component>

  <component name="ccpm-prd">
    <path>ccpm/skills/ccpm-prd/</path>
    <status>fail</status>
    <issues>
      <issue severity="critical">
        <location>SKILL.md:11</location>
        <description>Markdown heading inside XML tag: "## How PRD Management Works" inside essential_principles</description>
        <fix>Remove markdown headings; use semantic XML subtags or plain prose instead</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:15-43</location>
        <description>Multiple markdown headings (### PRD Lifecycle, ### File Structure, ### PRD Sections, ### Relationship to Epics) inside essential_principles tag</description>
        <fix>Convert to nested XML tags like prd_lifecycle, file_structure, prd_sections, relationship_to_epics</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:73</location>
        <description>Markdown heading "## Available Operations" inside operation_details tag</description>
        <fix>Remove heading; tag name already conveys purpose</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>workflows/.gitkeep</location>
        <description>Routing table references 5 commands (prd-new, prd-edit, prd-parse, prd-list, prd-status) but workflows/ directory contains only .gitkeep</description>
        <fix>Create workflow files for each command or remove routing references</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="info">
        <location>SKILL.md:17-20</location>
        <description>Bold markdown (**Creation**, **Editing**, etc.) used within XML for emphasis</description>
        <fix>Consider using plain prose; emphasis is acceptable but optional to convert</fix>
        <pre_identified>false</pre_identified>
      </issue>
    </issues>
  </component>

  <component name="ccpm-testing">
    <path>ccpm/skills/ccpm-testing/</path>
    <status>fail</status>
    <issues>
      <issue severity="critical">
        <location>SKILL.md:44-45</location>
        <description>Hardcoded paths reference wrong project and non-existent directories: /home/caio/Developer/Claude/ccpm-audit/ccpm/commands/testing/prime.md and run.md. These paths point to ccpm-audit (not ccpm) and commands/ (not workflows/).</description>
        <fix>Replace with relative paths: workflows/prime-testing.md and workflows/run-tests.md</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:1-110</location>
        <description>Missing required objective tag. Every skill must declare its purpose.</description>
        <fix>Add objective tag at the start of the body describing what the skill does and why.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:1-110</location>
        <description>Missing required quick_start tag. The intake tag is not a substitute.</description>
        <fix>Add quick_start tag with immediate actionable guidance for common use.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:7,11,15,19,51,79</location>
        <description>Markdown headings (## and ###) inside XML tags violates pure XML structure requirement. Found: "## How CCPM Testing Works", "### 1. Prime", "### 2. Run", etc.</description>
        <fix>Remove all markdown headings from within XML tags. Use nested XML elements or plain text with line breaks for structure.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>SKILL.md:44-45</location>
        <description>Routing table references files that don't exist at the specified paths. Actual files are at workflows/prime-testing.md and workflows/run-tests.md.</description>
        <fix>Update routing table to use correct relative paths to existing workflow files.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="warning">
        <location>workflows/prime-testing.md:9,14-16,73,118,296,341,382,418,426,461,495</location>
        <description>Workflow files use markdown headings instead of XML structure. While reference files are lower priority, they should also migrate to pure XML.</description>
        <fix>Convert markdown headings to semantic XML tags when refactoring.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="info">
        <location>workflows/prime-testing.md:504, workflows/run-tests.md:113</location>
        <description>Workflow files end with bare $ARGUMENTS placeholder without context on how it should be populated.</description>
        <fix>Wrap in XML tag or add documentation about argument handling.</fix>
        <pre_identified>false</pre_identified>
      </issue>
    </issues>
  </component>

  <component name="ccpm-worktree">
    <path>ccpm/skills/ccpm-worktree/</path>
    <status>fail</status>
    <issues>
      <issue severity="critical">
        <location>SKILL.md:1-254</location>
        <description>Missing required XML tags: objective and quick_start. Only success_criteria is present.</description>
        <fix>Add objective tag explaining what the skill does and why. Add quick_start tag with immediate actionable guidance.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:7</location>
        <description>Markdown heading "## How Worktrees Work in CCPM" inside XML tag violates pure XML structure requirement.</description>
        <fix>Remove markdown heading. Content can remain as prose within essential_principles or restructure into semantic XML subtags.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:11</location>
        <description>Markdown heading "### Key Concepts" inside XML tag.</description>
        <fix>Replace with XML subtag like key_concepts or use bold text (**Key Concepts:**).</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:24</location>
        <description>Markdown heading "### Parallel Agent Work" inside XML tag.</description>
        <fix>Replace with XML subtag or bold text.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:33</location>
        <description>Markdown heading "### Best Practices" inside XML tag.</description>
        <fix>Replace with XML subtag like best_practices or bold text.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>SKILL.md:69-102</location>
        <description>Workflow sections use "**Create Worktree Workflow**" as bold text which works, but inconsistent with other workflows that also use bold headers.</description>
        <fix>Consider using description subtags within each workflow for consistency.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="info">
        <location>SKILL.md:224-235</location>
        <description>References "worktree-manager" subagent but no verification this subagent exists in the codebase.</description>
        <fix>Verify worktree-manager subagent exists at ccpm/agents/worktree-manager.md or remove reference.</fix>
        <pre_identified>false</pre_identified>
      </issue>
    </issues>
  </component>

  <component name="ccpm-epic">
    <path>ccpm/skills/ccpm-epic/</path>
    <status>fail</status>
    <issues>
      <issue severity="critical">
        <location>SKILL.md:6-45</location>
        <description>Markdown headings inside XML tag: ## and ### headings within essential_principles violates pure XML structure</description>
        <fix>Convert markdown headings to nested semantic XML tags: epic_lifecycle, key_concepts, etc.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:91-116</location>
        <description>Markdown heading "## Available Epic Commands" inside command_reference tag</description>
        <fix>Remove markdown heading or restructure as command_reference content without ## prefix</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:118-156</location>
        <description>Markdown heading "## Common Epic Workflows" inside workflow_patterns tag</description>
        <fix>Remove markdown heading; the XML tag name already conveys section purpose</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:1-167</location>
        <description>Missing required objective tag - skill does not declare its purpose</description>
        <fix>Add objective tag near top describing: "Routes epic management requests to appropriate slash commands for planning, execution, tracking, and completion"</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:1-167</location>
        <description>Missing required quick_start tag - no immediate actionable guidance</description>
        <fix>Add quick_start with common entry point: "For new epic: /pm:prd-parse {name}, then /pm:epic-oneshot {name}"</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="warning">
        <location>references/epic-lifecycle.md:1-275</location>
        <description>Reference file uses markdown structure (## headings) instead of XML structure</description>
        <fix>Lower priority than SKILL.md but should migrate to pure XML for consistency</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>workflows/*.md</location>
        <description>All 8 workflow files use markdown ## headings instead of XML structure</description>
        <fix>Migrate to pure XML structure for consistency with skill standards</fix>
        <pre_identified>true</pre_identified>
      </issue>
    </issues>
  </component>

  <component name="ccpm-issue">
    <path>ccpm/skills/ccpm-issue/</path>
    <status>fail</status>
    <issues>
      <issue severity="critical">
        <location>workflows/:all</location>
        <description>Duplicate workflow files exist for every workflow. Both short names (analyze.md, start.md, close.md, sync.md, status.md, edit.md, reopen.md) AND long names (analyze-issue.md, start-issue.md, close-issue.md, sync-issue.md, status-issue.md, edit-issue.md, reopen-issue.md) exist. SKILL.md routes to short names but both versions contain similar content.</description>
        <fix>Delete the duplicate *-issue.md files. Keep the short-named versions (analyze.md, start.md, etc.) as they are referenced in SKILL.md routing table.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:6-23</location>
        <description>Markdown headings (## and ###) used inside essential_principles XML tag. Lines 7: "## How CCPM Issue Management Works", lines 9, 12, 15, 18, 21 contain "### Principle N:" headings. Pure XML structure required.</description>
        <fix>Convert markdown headings to semantic XML tags like principle name="local-first" or use numbered list format within the XML tag.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:55-69</location>
        <description>Markdown headings inside reference_index XML tag. Line 56: "## Shared Knowledge". Pure XML structure violated.</description>
        <fix>Remove "## Shared Knowledge" heading or convert to XML attribute on parent tag.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:71-84</location>
        <description>Markdown heading inside workflows_index XML tag. Line 72: "## Available Workflows".</description>
        <fix>Remove markdown heading; the XML tag name is self-documenting.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="critical">
        <location>SKILL.md:1-97</location>
        <description>Missing required XML tags: objective and quick_start are not present. These are mandatory for all skills.</description>
        <fix>Add objective describing what the skill does and quick_start with immediate actionable guidance before the intake section.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="critical">
        <location>workflows/analyze.md:11-12</location>
        <description>Markdown headings inside process XML tag: "## Step 1: Validate Issue Exists" and similar throughout. This pattern repeats in all short-named workflow files.</description>
        <fix>Convert "## Step N:" headings to step name="validate-issue" XML tags or numbered list format.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>references/frontmatter-operations.md:8,35,53,76,113,129,154</location>
        <description>Markdown headings inside XML tags throughout reference file. Example: Line 8 "## Task File Frontmatter" inside frontmatter_structure.</description>
        <fix>Remove markdown headings from inside XML tags; the tag names are descriptive enough.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>references/datetime-handling.md:13,37,62,103,129,156</location>
        <description>Markdown headings inside XML tags: "## Get Current Datetime", "## ISO 8601 Format", etc.</description>
        <fix>Remove markdown headings from inside XML tags.</fix>
        <pre_identified>true</pre_identified>
      </issue>
      <issue severity="warning">
        <location>references/parallel-streams.md:1-190</location>
        <description>Reference file uses only markdown headings without any XML structure. Should use pure XML for consistency.</description>
        <fix>Wrap content in semantic XML tags like overview, stream_patterns, parallelization_strategies.</fix>
        <pre_identified>false</pre_identified>
      </issue>
      <issue severity="info">
        <location>workflows/show.md</location>
        <description>File exists in glob results but not referenced in SKILL.md routing table (line 47 routes "show" to workflows/show.md but no corresponding *-issue.md duplicate).</description>
        <fix>Verify show.md workflow is complete and properly structured.</fix>
        <pre_identified>false</pre_identified>
      </issue>
    </issues>
  </component>

  <pre_identified_validation>
    <issue expected="ccpm-testing hardcoded path at lines 44-45" found="true">Confirmed: /home/caio/Developer/Claude/ccpm-audit/ccpm/commands/testing/ paths exist and are incorrect</issue>
    <issue expected="Markdown headings in XML across all skills" found="true">Confirmed: All 6 skills have ## and ### headings inside XML tags</issue>
    <issue expected="Duplicate workflows in ccpm-issue" found="true">Confirmed: 7 pairs of duplicate workflow files (short-name.md and long-name-issue.md)</issue>
  </pre_identified_validation>

  <metadata>
    <audit_date>2024-12-21</audit_date>
    <auditor>skill-auditor subagent (6 parallel instances)</auditor>
    <confidence level="high">Based on automated file analysis and pattern matching against skill best practices</confidence>
  </metadata>
</audit_results>
```
