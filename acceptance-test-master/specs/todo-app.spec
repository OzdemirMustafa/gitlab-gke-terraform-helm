# Acceptance Test
## Todo App


* Goto ToDo list page

* Given Empty ToDo list
* When I write "buy some milk" to text box and click to add button
* Then I should see "buy some milk" item in ToDo list

* Given ToDo list with "buy some milk" item
* When I write "enjoy the assignment" to text box and click to add button
* Then I should see "enjoy the assignment" insterted to ToDo list below "buy some milk"

* Given ToDo list with "buy some milk" item
* When I click on "buy some milk" text
* Then I should see "buy some milk" item marked as "buy some milk"

* Given ToDo list with marked "buy some milk" item
* When I click on "buy some milk" text
* Then I should see mark of "buy some milk" item should be cleared as "buy some milk"

* Given ToDo list with "enjoy the assignment" item
* When I click on delete button next to "enjoy the assignment" items
* Then List should be empty

* Given ToDo list with "enjoy the assignment" and "drink water" item in order
* When I click on delete button next to "rest for a while" item
* Then I should see "drink water" item in ToDo list