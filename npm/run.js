#!/usr/bin/env node
'use strict';

const { spawnSync } = require('child_process');
const path = require('path');
const fs = require('fs');

const ext = process.platform === 'win32' ? '.exe' : '';
const binary = path.join(__dirname, 'bin', `datpaq${ext}`);

if (!fs.existsSync(binary)) {
  console.error('datpaq binary not found. Try reinstalling: npm install -g datpaq');
  process.exit(1);
}

const result = spawnSync(binary, process.argv.slice(2), { stdio: 'inherit' });
process.exit(result.status ?? 1);
