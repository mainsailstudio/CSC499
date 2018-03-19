/*
	Title:	Combine perms service
	Author:	Connor Peters
	Date:	2/24/2018
	Desc:	Simply combines the locks and keys into one string and returns them in an array
*/

import { Injectable } from '@angular/core';

@Injectable()
export class CombinePermsService {

    constructor() { }

    // combinePerms - to correctly concat 2 slices of permutations into 1.
    // Needs 2 slices, 1 of locks and 1 of keys. It is assumed that they match up perfectly and will result in logic errors if they do not.
    combinePerms(locks: string[], keys: string[]): string[] {
        const combined = []; // assumes locks and keys are at the same index (SHOULD ALWAYS BE)
        for (let i = 0; i < locks.length; i++) {
            const combineString = locks[i] + keys[i];
            combined.push(combineString);
        }
        return combined;
    }

}

