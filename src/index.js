const schedule = require('node-schedule');
const { exec } = require('child_process');
const fs = require('fs');
const path = require('path');

const rule = process.env.SCHEDULE_RULE || '*/1 * * * *';
const sshRemoteUrl = process.env.SSH_REMOTE_URL || '';
const gitUserEmail = process.env.GIT_UESR_EMAIL || '';

if (!sshRemoteUrl || !gitUserEmail) {
  console.log('sshRemoteUrl and gitUserEmail can not be empty.');
  process.exit(1);
}

console.log(`=== schedule commit start, commit schedule rule: ${rule}===`);

function gitInit() {
  const command = 'git init';
  console.log('initialize git repo');
  exec(command, (err, stdout, stderr) => {
    if (err) {
      console.log('gitInit error: ', err);
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
  });
}

function setGitConfig() {
  const command = `git config --local user.email ${gitUserEmail}`;
  console.log('initialize git repo');
  exec(command, (err, stdout, stderr) => {
    if (err) {
      console.log('gitInit error: ', err);
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
  });
}

function setRemoteUrl() {
  const command = `git remote set-url origin ${sshRemoteUrl}`;
  console.log(`set remote url to ${sshRemoteUrl}`);
  exec(command, (err, stdout, stderr) => {
    if (err) {
      console.log('setRemoteUrl error: ', err);
    }
  });
}

function gitpush(callback) {
  const command = "git add . && git commit -m 'commit' && git push origin master";
  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.log(`gitpush error: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.log(`stderr: ${stderr}`);
  });
}

function renameSync() {
  const pathA = path.resolve(__dirname, './test.txt');
  const pathB = path.resolve(__dirname, './test-1.txt');
  if (fs.existsSync(pathA)) {
    fs.renameSync(pathA, pathB);
  } else {
    fs.renameSync(pathB, pathA);
  }
}

function main() {
  gitInit();
  setGitConfig();
  setRemoteUrl();

  const j = schedule.scheduleJob(rule, () => {
    console.log('start to commit, next commit date: ', j.nextInvocation());
    renameSync();
    gitpush();
  });

  console.log(j.nextInvocation());
}

main();
