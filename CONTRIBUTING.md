# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. 

Please note we have a code of conduct, please follow it in all your interactions with the project.

## Fork Process

1. Ensure that you've installed the Golang (minimum 1.13) in your system.
2. For this project into your own Github account.
3. Clone the `grule-rule-engine` forked repository on your account.
4. Enter the cloned directory.
5. Apply new "upstream" to original `hyperjumptech/grule-rule-engine` git 
4. Now you can work on your account
5. Remember to pull from your upstream often. `git pull upstream master`

## Pull Request Process

1. Make sure you always have the most recent update from your upstream. `git pull upstream master`
2. Resolve all conflict, if any.
3. Make sure `make test` always successful (you wont be able to create pull request if this fail, circle-ci, travis-ci and azure-devops will make sure of this.)
4. Push your code to your project's master repository.
5. Create PullRequest. 
    * Go to `hithub.com/hyperjumptech/grule-ruleengine`
    * Select `Pull Request` tab
    * Click "New pull request" button
    * Click "compare across fork"
    * Change the source head repository from your fork and target is `hyperjumptech/grule-rule-engine`
    * Hit the "Create pull request" button
    * Fill in all necessary information to help us understand about your pull request.  

