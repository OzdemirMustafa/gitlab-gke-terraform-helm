const {openBrowser, write, click, closeBrowser, goto, text, focus, $} = require('taiko');
const assert = require("assert").strict;
const headless = process.env.headless_chrome.toLowerCase() === 'true';

beforeSuite(async () => {
    await openBrowser({headless: headless,
        args: [
            '--disable-gpu',
            '--disable-dev-shm-usage',
            '--disable-setuid-sandbox',
            '--no-first-run',
            '--no-sandbox',
            '--no-zygote',
            '--lang=tr'
        ],})
});

afterSuite(async () => {
    await closeBrowser()
});


step("Goto ToDo list page", async () => {
    await goto('http://35.225.7.155/')
});

step("Given Empty ToDo list", async () => {
    assert.ok(!await text("buy some milk").exists())
});

step("When I write <text> to text box and click to add button", async (text) => {
    await focus($('#todoText'))
    await write(text);
    await click($('#addButton'))
});

step("Then I should see <textMilk> item in ToDo list", async (textMilk) => {
    assert.ok(await text(textMilk).exists())
});


step("Given ToDo list with <text> item", async (textMilk) => {
    assert.ok(await text(textMilk).exists())
});

step("Then I should see <secondItem> insterted to ToDo list below <firstItem>", async (secondItem, firstItem) => {
    assert.ok(await text(firstItem).exists())
    assert.ok(await text(secondItem).exists())
});

step("When I click on <textMilk> text", async (textMilk) => {
    await click(textMilk)
});

step("Then I should see <textMilk> item marked as <textMilk>", async (textMilk) => {
    assert.ok(await text(textMilk).exists())
});

step("Given ToDo list with marked <textMilk> item", async (textMilk) => {
    assert.ok(await text(textMilk).exists())
});

step("Then I should see mark of <textMilk> item should be cleared as <textMilk>", async (textMilk) => {
await click(textMilk)
});


step("When I click on delete button next to <restWhile> items", async (restWhile) => {
    await click( text("x"))
    assert.ok(await text(restWhile).exists())

    await click( text("x"))
});


step("Then List should be empty", async () => {
    assert.ok(!await $(".todoList").exists())
});


step("Given ToDo list with <textOne> and <textTwo> item in order", async (textOne, textTwo) => {
    await focus($('#todoText'))
    await write(textOne);
    await click($('#addButton'))

    await focus($('#todoText'))
    await write(textTwo);
    await click($('#addButton'))
});

step("When I click on delete button next to <texts> item", async (texts) => {
    await click(text("x"))
});
















