# Contributors

## Definitions

* :version = Semantic versioning version string (ex. 2.1.15).
* :name = Name of a version-control user.
* :requirement-number = Requirement number.
* :requirement = Requirement content.
* :issue-id = ID for an issue.

## Repository Layout

```
README.md (brief project description and installation and running instructions)
CONTRIBUTORS.md (instructions for contributors)
Makefile (builds everything, minimally has 'doc' target that converts 'doc'
          markdowns to PDFs)
?LICENSE.txt (optional license)
doc/
    requirements.md (release requirements)
    design.md (release design)
    requirements.pdf (PDF generated from 'requirements.md')
    design.pdf (PDF generated from 'design.md')
```

* 'requirements.md' has sections labelled with header ':version Requirements'
  and each section contains a numbered list of requirements as incomplete
  sentences.
* 'design.md' has sections labelled with header ':version Design' and each
  section contains sufficient materials with relevant requirements cited
  parenthetically.

## Branching

* Repository work is done on branch 'repo'.
* Requirements work is done on branch ':version-requirements'.
* Design work is done on branch ':version-design'.
* Users can have a playground branch called ':name-playground' that can never be
  merged into another branch.
* Other branches are named after their issue ID.
* All branches, except 'repo', are branched from branch ':version'.
* All work, except as part of branch 'repo', must be done as part of a release.

## Process:

1. `git checkout master`
2. `git checkout -b :version`
3. Create issue 'Add :version Requirements' labelled ':version'.
4. Assign issue.
5. `git checkout -b :version-requirements`
6. Add version requirements to 'doc/requirements.md'.
7. Commit and push.
8. Create a pull-request for `:version-requirements`.
9. Add issues labelled ':version' and 'requirement' named with
   ':version :requirement-number. :requirement'.
10. Create labels named ':version-:requirement-label' for each requirement.
10. Repeat steps 3 through 8 for design.
11. Break the design into appropriate tasks labelled with ':version' and the
    requirements they address.
12. `git checkout -b :version-:issue-id` from branch `:version` and repeat
    workflow up to pull-request.
13. Close requirement issues once their last issue is merged.
14. Create a pull-request for ':version' branch once all requirement issues are
    closed.
15. Tag master with a release with version ':version'.
