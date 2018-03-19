/*
	Title:	Permutations service
	Author:	Connor Peters
	Date:	3/12/2018
    Desc:	To generate a limited number of permutations (i.e. 1234, 5678, 78910 from 10P4) on the client side
            This is a javscript re-write of the og permutation.go file on the back-end developed for testing
	Usage:  Call LimPerms with a slice of strings to permutate
*/

import { Injectable } from '@angular/core';

@Injectable()
export class PermutateService {

  constructor() { }

    // Generates the subsets of the passed array, and then permutes each subset
    generateLimPerms (toPerm: string[], num: number) {
        const subsets = this.getSubsets(toPerm, num);
        const perms = this.getPerms(subsets);
        return perms;
    }

    // Generates each unique subset of the array that is passed in, with the num being the limiting factor
    getSubsets(locks: string[], num: number): string[][] {
        const res = [];
        const findSubset = [];

        // helper 'fat arrow' lambda function
        const helper = (arr: string[], subset: string[], n: number) => {
            if (subset.length === num) {
                const tmp = subset.slice();
                res.push(tmp);
                return;
            }

            if (n === arr.length) {
                return;
            }

            subset.push(arr[n]);
            // recursion starts here
            helper(locks, subset, n + 1);
            subset.pop();
            helper(locks, subset, n + 1);
        };

        helper(locks, findSubset, findSubset.length);
        return res;
    }

    // A javascript implementation of Heap's algorithm
    heapPermutation(arr: string[]): string[] {
        const res = [];

        const helper = (array: string[], n: number) => {
            if (n === 1) {
                const tmp = array.join('').slice();
                res.push(tmp);
            } else {
                for (let i = 0; i < n; i++) {
                    // recursion people
                    helper(array, n - 1);
                    if (n % 2 === 1) {
                        const tmp = array[i];
                        array[i] = array[n - 1];
                        array[n - 1] = tmp;
                    } else {
                        const tmp = array[0];
                        array[0] = array[n - 1];
                        array[n - 1] = tmp;
                    }
                }
            }
        };
        helper(arr, arr.length);
        return res;
    }

    // This is the organizer for heap's algo
    // It takes a 2d array of each subset (generated above) and returns a single array
    // With each of heap's permutations as a single string with no delimiter (joined)
    getPerms(perms: string[][]): string[] {
        const res = [];
        for (let i = 0; i < perms.length; i++) {
            const list = perms[i].slice();
            const tmp = this.heapPermutation(list).slice();
            for (let j = 0; j < tmp.length; j++) {
                res.push(tmp[j]);
            }
        }
        return res;
    }

}

