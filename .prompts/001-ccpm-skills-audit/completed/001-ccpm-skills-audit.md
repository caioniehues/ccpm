<objective>
Execute comprehensive audits of all 6 CCPM skills in parallel using the /taches:audit-skill command.

Purpose: Gather structured findings about skill compliance, XML structure, and routing issues
Output: skills-audit.md with XML-structured findings per skill
</objective>

<context>
Project: CCPM (Claude Code Project Manager)
Location: ccpm/skills/

Pre-identified critical issues to validate:
- ccpm-testing/SKILL.md:44-45: Hardcoded wrong path (/home/caio/Developer/Claude/ccpm-audit/...)
- All 6 skills: Markdown headings (##) inside XML tags (violates pure XML structure)
- ccpm-issue: Duplicate workflow files (e.g., analyze.md AND analyze-issue.md)
</context>

<requirements>
Execute these 6 audits IN PARALLEL using Task tool with skill-auditor subagent:
1. /taches:audit-skill ccpm/skills/ccpm-context
2. /taches:audit-skill ccpm/skills/ccpm-prd
3. /taches:audit-skill ccpm/skills/ccpm-testing
4. /taches:audit-skill ccpm/skills/ccpm-worktree
5. /taches:audit-skill ccpm/skills/ccpm-epic
6. /taches:audit-skill ccpm/skills/ccpm-issue

For maximum efficiency, spawn ALL 6 Task agents in a SINGLE message with multiple tool calls.
</requirements>

<output_structure>
Save to: .prompts/001-ccpm-skills-audit/skills-audit.md

Write findings INCREMENTALLY using streaming writes:
1. Create file with XML skeleton
2. Append each skill's findings as its audit completes
3. Finalize summary counts at end

Use this XML structure:

```xml
<audit_results type="skills">
  <summary>
    <total_components>6</total_components>
    <critical_issues>N</critical_issues>
    <warning_issues>N</warning_issues>
    <info_issues>N</info_issues>
  </summary>

  <component name="{skill-name}">
    <path>ccpm/skills/{skill-name}/</path>
    <status>pass|fail</status>
    <issues>
      <issue severity="critical|warning|info">
        <location>file:line</location>
        <description>What's wrong</description>
        <fix>How to fix it</fix>
        <pre_identified>true|false</pre_identified>
      </issue>
    </issues>
  </component>

  <pre_identified_validation>
    <issue expected="ccpm-testing hardcoded path" found="true|false"/>
    <issue expected="Markdown in XML" found="true|false"/>
    <issue expected="Duplicate workflows" found="true|false"/>
  </pre_identified_validation>

  <metadata>
    <confidence level="high">Based on automated audit tools</confidence>
  </metadata>
</audit_results>
```
</output_structure>

<summary_requirements>
Create .prompts/001-ccpm-skills-audit/SUMMARY.md

One-liner format: "{N} critical, {N} warning, {N} info issues across 6 CCPM skills"

Key Findings:
- List top 3 most impactful issues
- Note if pre-identified issues were confirmed

Next Step: Run 002-ccpm-commands-audit
</summary_requirements>

<success_criteria>
- All 6 skills audited
- Structured XML output created
- Pre-identified issues validated
- SUMMARY.md with substantive one-liner
- Ready for downstream consumption by 004 and 005
</success_criteria>
