<objective>
Execute comprehensive audits of all 10 CCPM subagents in parallel.

Purpose: Gather structured findings about agent configuration, tool permissions, and prompt quality
Output: agents-audit.md with XML-structured findings per agent
</objective>

<context>
Project: CCPM (Claude Code Project Manager)
Location: ccpm/agents/

Previous audits:
- @.prompts/001-ccpm-skills-audit/skills-audit.md
- @.prompts/002-ccpm-commands-audit/commands-audit.md

Pre-identified critical issues to validate:
- test-runner.md: MISSING Bash tool (needs it to run tests!)
- code-analyzer.md, file-analyzer.md, parallel-worker.md, test-runner.md: Markdown headings in body
- code-analyzer.md, file-analyzer.md, test-runner.md: Over-permissioned (have Task/Agent but are leaf nodes)
</context>

<requirements>
Execute ALL 10 audits IN PARALLEL using a single message with multiple Task tool calls:

1. /taches:audit-subagent ccpm/agents/code-analyzer.md
2. /taches:audit-subagent ccpm/agents/epic-planner.md
3. /taches:audit-subagent ccpm/agents/file-analyzer.md
4. /taches:audit-subagent ccpm/agents/github-syncer.md
5. /taches:audit-subagent ccpm/agents/parallel-orchestrator.md
6. /taches:audit-subagent ccpm/agents/parallel-worker.md
7. /taches:audit-subagent ccpm/agents/prd-architect.md
8. /taches:audit-subagent ccpm/agents/task-decomposer.md
9. /taches:audit-subagent ccpm/agents/test-runner.md
10. /taches:audit-subagent ccpm/agents/worktree-manager.md

Pay special attention to:
- Tool permission analysis (what tools each agent has)
- Leaf node detection (agents that shouldn't spawn other agents)
- Missing required tools
</requirements>

<output_structure>
Save to: .prompts/003-ccpm-agents-audit/agents-audit.md

Write incrementally as each audit completes.

```xml
<audit_results type="agents">
  <summary>
    <total_components>10</total_components>
    <critical_issues>N</critical_issues>
    <warning_issues>N</warning_issues>
    <info_issues>N</info_issues>
  </summary>

  <component name="{agent-name}">
    <path>ccpm/agents/{agent-name}.md</path>
    <status>pass|fail</status>
    <tools_declared>[list of tools]</tools_declared>
    <is_leaf_node>true|false</is_leaf_node>
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
    <issue expected="test-runner missing Bash" found="true|false"/>
    <issue expected="Markdown headings in body" found="true|false"/>
    <issue expected="Over-permissioned leaf nodes" found="true|false"/>
  </pre_identified_validation>

  <tool_permission_analysis>
    <agent name="{name}" tools="[list]" should_have="[list]" should_remove="[list]"/>
  </tool_permission_analysis>

  <metadata>
    <confidence level="high">Based on automated audit tools</confidence>
  </metadata>
</audit_results>
```
</output_structure>

<summary_requirements>
Create .prompts/003-ccpm-agents-audit/SUMMARY.md

One-liner: "{N} critical, {N} warning, {N} info issues across 10 CCPM agents"

Key Findings:
- Critical missing tools (blocking issues)
- Over-permissioned agents
- Structural violations

Next Step: Run 004-ccpm-integration-research
</summary_requirements>

<success_criteria>
- All 10 agents audited
- Pre-identified issues validated
- Tool permission analysis complete
- Ready for integration research
</success_criteria>
