/*
	Title:	Hash SHA256 service
	Author:	Connor Peters
	Date:	2/24/2018
	Desc:	Takes in a "toHash" slice and returns the corresponding hashes
	NOTE:	This was implemented for sake of speed, Bcrypt is a better password hasher but this isn't a password baby
*/

import { Injectable } from '@angular/core';
import * as shajs from 'sha.js';

@Injectable()
export class HashSha256Service {

  constructor() { }

  // takes in the 'toHash' array, hashes each item and then returns a new array with the hashes
  hashPermsSHA256(toHash: string[]): string[] {
    const hashed = [];
    for (let i = 0; i < toHash.length; i++) {
        const hashString = shajs('sha256').update(toHash[i]).digest('hex');
        hashed.push(hashString);
    }
    return hashed;
  }

}

