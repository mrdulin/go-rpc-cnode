const schedule = require('node-schedule');
const { exec } = require('child_process');
const fs = require('fs');
const path = require('path');

const rule = '3 * * * * *';
const j = schedule.scheduleJob(rule, () => {
  console.log('start to commit');
  modifyfile(() => {
    gitpush();
  });
});

const command = "git add . && git commit -m 'commit' && git push origin master";
function gitpush(callback) {
  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
  });
}

// gitpush();

function modifyfile(callback) {
  const filepath = path.resolve(__dirname, './test.txt');
  fs.open(filepath, 'w+', (err, fd) => {
    if (err) return console.error(err);
    const data = '\n';
    fs.write(fd, data, (error, bytesWritten, buffer) => {
      if (error) return console.error(error);
      callback && callback();
    });
  });
}

// modifyfile();
