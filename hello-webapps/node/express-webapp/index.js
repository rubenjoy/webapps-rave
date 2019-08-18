const express = require('express');
const app = new express();

const callback = (req,res) => {
    res.send('hello world');
}

app.get('/', callback);
app.listen(3000, () => console.log('App listening on port 3000!'));