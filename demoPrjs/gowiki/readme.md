Writing Web Applications


https://go.dev/doc/articles/wiki/


Introduction
Covered in this tutorial:

Creating a data structure with load and save methods
Using the net/http package to build web applications
Using the html/template package to process HTML templates
Using the regexp package to validate user input
Using closures
Assumed knowledge:

Programming experience
Understanding of basic web technologies (HTTP, HTML)
Some UNIX/DOS command-line knowledge


Other tasks
Here are some simple tasks you might want to tackle on your own:

X Store templates in tmpl/ and page data in data/.
X Add a handler to make the web root redirect to /view/FrontPage.
Spruce up the page templates by making them valid HTML and adding some CSS rules.
Implement inter-page linking by converting instances of [PageName] to
<a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)