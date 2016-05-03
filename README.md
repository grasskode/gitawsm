# gitawsm

gitawsm is an awesome extension to git to make branch management easier. A typical use case would be in a micro service deployment where one feature spans across multiple projects. A developer working on multiple projects can manage a feature branch using gitawsm.

## Installation

Install using <b>go get</b>

<code>go get github.com/grasskode/gitawsm</code>

## Usage

A typical use case would be as follows :

### Create a new branch
Execute <code>gitawsm branch</code> from anywhere. Let's say we need to work on a new analytics feature.

<code>gitawsm branch feature/analytics</code>

This will create a new branch by the name of feature/analytics. Any branch name provided here should be a valid git ref name.

### Add projects to branch
Let's say that the branch affects <code>webapp</code> and <code>api</code> projects. We can add the projects to the gitawsm branch using <code>gitawsm add</code>. This can be done in two ways.


One, we switch to the project and add it to the branch.

<code>cd /path/to/project/webapp</code>
<code>gitawsm add feature/analytics</code>

Or, we issue the command globally with project paths.

<code>gitawsm add feature/analytics /path/to/project/webapp /path/to/project/api</code>

### List branches and projects
At any point of time you can check the branches and associated projects using <code>gitawsm list</code>

<code>gitawsm list</code> will list all branches with associated projects.

<code>gitawsm list "feature.*"</code> will list all branches matching "feature.*" regexp. The regexp should be a valid re2 expression.

### Checkout branch
Checkout the gitawsm branch using <code>gitawsm checkout</code>. This will checkout the branch in all projects associated with the branch.

<code>gitawsm checkout feature/analytics</code>

### Push to remote
Work across projects and once you are ready to push your changes, use <code>gitawsm push</code> to push the branch across all projects.

<code>gitawsm push feature/analytics</code>

### Pull from upstream
Update all projects for changes in the upstream using <code>gitawsm pull</code>.

<code>gitawsm pull feature/anaytics</code>
