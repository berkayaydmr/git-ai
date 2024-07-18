# Git AI
This is a simple project that uses GPT to get a review of the difference between two branches of git repository.

## Requirements
1. Golang programming language
2. Build go project with the following command
    ```
    go build -o git-ai .
    ```
    then move the binary and layout to the bin directory with the following command
    ```
    sudo mv git-ai /usr/local/bin
    sudo mv review-layout.txt /usr/local/bin
    ```
    after that make alias for the binary with the following command
    ```
    alias git-ai='/usr/local/bin/git-ai'
    ```
    or you can use makefile for this `make all` command.

3. GPT API key and gpt engine version you can save(encrypted) your api key and gpt version  with the following command
    ```
    git ai set key NAME_FOR_KEY API_KEY
    ```

4. `CRYPTOGRAPHER_KEY` environment variable for encryption and decryption of the api key should be set
    you can set env value with the following command in diffrent operating systems

    **Linux/Macos**
    ```
    export CRYPTOGRAPHER_KEY=your-32-byte-key-here
    ```

    **Windows**
    ```
    set CRYPTOGRAPHER_KEY=your-32-byte-key-here
    ```

    ***Windows Powershell***
    ```
    $env:CRYPTOGRAPHER_KEY="your-32-byte-key-here"
    ```

## How to use
1. Clone the repository

3. Run the following command in project directory

    ```
    git ai diff /path/to/repo BRANCH1 BRANCH2 
    ```
    or in the project directory
    ```
    git ai diff BRANCH1 BRANCH2 
    ```
    
    you may prefer to use makefile for this command.

4. The result will be shown in the terminal

## Commands
- `git ai diff /path/to/repo branch1 branch2` : This command is used to get the difference between two branches of a git repository
- `git ai set key nameforkey` : This command is used to save the api key and gpt version for future use
- `git ai remove key nameforkey` : This command is used to remove the saved api key and gpt version

## Example
inputs are pointed with `->`

```
-> git ai diff ~/path/to/repo main init
Repository URL:  ~/path/to/repo
Select an API key:
There is no API key, please add one or use custom key option to enter a custom key manually select 0
-> 0
Enter the API key: -> sk-proj-vKve7ZBSTxjlse8Fd28BT3BlbkFJQRFjt0t8GNGJ0jcmxkX2
1. gpt-3.5-turbo
2. gpt-4
3. gpt-3
4. not-settled
-> 1

Waiting for GPT response.

**language-learning-api**

From e026003b638fb1f45a51dd1b3333914f7faf38aa to 9e55503d2f2342f75a51ed81a2941e5199d615e6 (21 commits)
10 files changed, 1091 insertions(+), 0 deletions(-)

The following changes were made:
- Added Dockerfile with 13 insertions and 0 deletions
- Added main.go file for dbseeder with 34 insertions and 0 deletions
- Added main.go file for server with 81 insertions and 0 deletions
- Added docker-compose.yaml with 9 insertions and 0 deletions
- Added go.mod with 21 insertions and 0 deletions
- Added go.sum with 49 insertions and 0 deletions
- Added makefile with 10 insertions and 0 deletions
- Added error.go with 38 insertions and 0 deletions
- Added SQL files for storage with various insertions and 0 deletions
- Added middleware files with various insertions and 0 deletions
- Added transport files with various insertions and 0 deletions
- Added handler.go with 27 insertions and 0 deletions
- Added utils.go with 26 insertions and 0 deletions

Review:
- The changes include the addition of various essential files for setting up a language learning API, such as Dockerfile, main.go files for dbseeder and server, docker-compose.yaml, makefile, error handling,
SQL files for storage, middleware, transport, and utility functions.
- The commits seem to cover a wide range of functionalities required for the API to function properly.
- The insertions indicate that a significant amount of code was added, which suggests thorough development work.
- It would be beneficial to review the SQL queries and ensure they are optimized and secure.
- Overall, the changes appear to lay a solid foundation for the language learning API.

Feedback:
- Great work on setting up the initial structure and functionality for the language learning API.
- Ensure to conduct thorough testing to validate the newly added features and functionalities.
- Consider adding documentation to explain the purpose and usage of each file and function.
- Keep up the good progress and continue to refine and enhance the API based on feedback and testing results.
```