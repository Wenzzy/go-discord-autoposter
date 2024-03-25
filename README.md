# go-discord-autoposter

### Getting access token from Discord

1. Open discord from your browser and login to your account. (or enable debug mode in discord app)
2. Press `Ctrl + Shift + I` to open developer tools. (or `Option + Command + I` on macOS)
3. Paste the following code in the console and press enter.

```js
const i = document.createElement("iframe");
i.onload = () =>
  console.log(
    "%c" + i.contentWindow.localStorage.getItem("token").slice(1, -1),
    "color: red;font-size: 24px"
  );
i.src = "about:blank";
document.body.appendChild(i);
```

4. You will see a red text in the console. That is your access token.

### Getting channel id from Discord

1. Open discord from your browser and login to your account. (or discord app)
2. Enable developer mode in discord settings.
3. Right click on the channel you want to post and click on `Copy ID`.
