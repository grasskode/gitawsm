# gitawsm

gitawsm is an awesome extension to git to make branch management easier. A typical use case would be in a micro service deployment where one feature spans across multiple projects. A developer working on multiple projects can manage a feature branch using gitawsm.

## Installation

You will need [golang](https://golang.org/) to install **gitawsm**.
Install golang using the instructions provided [here](https://golang.org/doc/install). 

Get gitawsm using <b>go get</b>

	go get github.com/grasskode/gitawsm

Install

	go install github.com/grasskode/gitawsm

The binary will be generated at **$GOPATH/bin**. You can add **$GOPATH/bin** to your **$PATH** to access gitawsm globally.

## Usage

A typical use case would be as follows :

### Create a new branch
Execute <code>gitawsm branch</code> from anywhere. Let's say we need to work on a new analytics feature.

	gitawsm branch feature/analytics

This will create a new branch by the name of feature/analytics. Any branch name provided here should be a valid git ref name.

### Add projects to branch
Let's say that the branch affects <code>webapp</code> and <code>api</code> projects. We can add the projects to the gitawsm branch using <code>gitawsm add</code>. This can be done in two ways.


One, we switch to the project and add it to the branch.

	cd /path/to/project/webapp
	gitawsm add feature/analytics

Or, we issue the command globally with project paths.

	gitawsm add feature/analytics /path/to/project/webapp /path/to/project/api

### List branches and projects
At any point of time you can check the branches and associated projects using <code>gitawsm list</code>

<code>gitawsm list</code> will list all branches with associated projects.

<code>gitawsm list "feature.*"</code> will list all branches matching "feature.*" regexp. The regexp should be a valid re2 expression.

### Checkout branch
Checkout the gitawsm branch using <code>gitawsm checkout</code>. This will checkout the branch in all projects associated with the branch.

	gitawsm checkout feature/analytics

### Push to remote
Work across projects and once you are ready to push your changes, use <code>gitawsm push</code> to push the branch across all projects.

	gitawsm push feature/analytics

### Pull from upstream
Update all projects for changes in the upstream using <code>gitawsm pull</code>.

	gitawsm pull feature/anaytics
