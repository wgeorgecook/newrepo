Small CLI I made to help me organize work I'm doing based on my Jira tickets

NewRepo clones a repo given by -r, to an _absolute_ path given by -p, and creates a new branch given by -n (optionally forked from branch given by -c)

Example:
newrepo -r "https://github.com/wgeorgecook/newrepo" -p "/Users/yourusername/Documents/" -n "New Test" -c "main"