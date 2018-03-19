import { Component, OnInit } from '@angular/core';
import { RegisterUserService } from '../register/register-user.service';
import { trigger, transition, style, animate, state } from '@angular/animations';

import { DashboardComponent } from './dashboard.component';

import { PermutateService } from '../hash/perm.service';
import { HashSha256Service } from '../hash/hash-sha256.service';
import { CombinePermsService } from '../hash/combine.service';
import { InitAccountService } from './init-account.service';
import { WordArray } from '../api/api.constants';

import * as shajs from 'sha.js';

export interface InitUser {
  id: string;
  email: string;
  fname: string;
  lname: string;
  securityLv: string;
}

@Component({
  selector: 'app-dashboard-test',
  templateUrl: './dashboard-test.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardTestComponent extends DashboardComponent implements OnInit {

  // get token data
  userDataJSON = localStorage.getItem('currentUser');
  // parse that data
  userData = JSON.parse(this.userDataJSON);
  // individual variable from the token data
  email = this.userData['email'];
  userID = this.userData['id'];
  jwToken = this.userData['token'];
  testLevel = this.userData['testLevel'];
  init = this.userData['init'];

  // auth array initialization, for now this works but it should grab from the database
  // eventually to preserve the users locks if they weren't numbers
  auths = [];

  // will also have to grab from db eventually
  displayLength = 4; // for now!!

  // the random words that autofill when creating a new setup
  randomWordArray = [];

  // this is the success variable that shows after the user account is configured
  showSuccess = false;
  showFail = true;

  constructor(
    private permService: PermutateService,
    private hashService: HashSha256Service,
    private combineService: CombinePermsService,
    private postConfigFormService: InitAccountService
  ) {
      super(); // extend the dashboard
    }

  validateRegistration(pass: string, confirmPass: string): boolean {
    if (pass === confirmPass) {
      return true;
    } else {
      return false;
    }
  }

  ngOnInit() {
    if (this.testLevel === 1) {
        console.log('Password here');
    } else if (this.testLevel === 2) {
        this.auths = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
        for (let i = 0; i < this.auths.length; i++) {
          this.randomWordArray.push(WordArray[Math.floor(Math.random() * 2627)]);
        }
        console.log('Easy security here');
    } else if (this.testLevel === 3) {
        this.auths = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
        for (let i = 0; i < this.auths.length; i++) {
          this.randomWordArray.push(WordArray[Math.floor(Math.random() * 2627)]);
        }
        console.log('Hard security here');
    }
  }

  configPass(formData: any) {
    // verify the pass and confirm pass are the same again
    const password = formData.value['tempPass'];
    const confirmPass = formData.value['confirmPass'];

    if (password !== confirmPass) {
      console.log('Passwords didnt match');
      return;
    } else {
      const hashedPass = shajs('sha256').update(password).digest('hex');
      this.postConfigFormService.postPassword(this.userID, hashedPass, this.jwToken).subscribe(
        suc => {
          console.log(suc);
        },
        err => {
          console.log(err );
        }
      );
    }
  }

  configAuths(formData: any) {
    const lockArray = [];
    const keyArray = [];

    // get all the locks and keys
    for (let i = 1; i <= this.auths.length; i++) {
      lockArray.push(this.auths[i - 1].toString());
      keyArray.push(formData.value['key' + i]);
    }

    // permutate the locks and keys
    // const lockPermArray = this.permService.generateLimPerms(lockArray, this.displayLength);
    const keyPermArray = this.permService.generateLimPerms(keyArray, this.displayLength);

    // combine the lock and key permutations
    // const combineArray = this.combineService.combinePerms(lockPermArray, keyPermArray);

    // hash the lock and key combos
    const hashArray = this.hashService.hashPermsSHA256(keyPermArray);

    this.postConfigFormService.postAuthArray(this.userID, lockArray, hashArray, this.jwToken).subscribe(
      suc => {
        console.log(suc);
        console.log('That was success, swapping to practice');
        this.swapDashboardComponent('practice');
      },
      err => {
        console.log(err );
      }
    );

  }

}
