import { Component, OnInit } from '@angular/core';
import { RegisterUserService } from '../register/register-user.service';
import { trigger, transition, style, animate, state } from '@angular/animations';

import { DashboardComponent } from './dashboard.component';

import { PermutateService } from '../hash/perm.service';
import { HashSha256Service } from '../hash/hash-sha256.service';
import { CombinePermsService } from '../hash/combine.service';
import { InitAccountService } from './init-account.service';
import { WordArray } from '../api/api.constants';
import { ActivityLogService } from '../activity-log/activity-log.service';

import { ConfigActivity } from '../activity-log/log.interface';

import * as shajs from 'sha.js';
import { UserConstantsService } from './user-constants/user-constants.service';

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
export class DashboardTestComponent implements OnInit {

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
  showInsertion = false;
  showSuccess = false;
  showFail = false;

  // activity logging variables
  loginStartTime = new Date().getTime();

  constructor(
    private permService: PermutateService,
    private hashService: HashSha256Service,
    private combineService: CombinePermsService,
    private postConfigFormService: InitAccountService,
    private activityLog: ActivityLogService
  ) { }

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
      this.showInsertion = true;
      const hashedPass = shajs('sha256').update(password).digest('hex');
      this.postConfigFormService.postPassword(this.userID, hashedPass, this.jwToken).subscribe(
        suc => {
          console.log(suc);
          this.showInsertion = false;
          this.showSuccess = true;
          this.showFail = false;
        },
        err => {
          console.log(err );
          this.showInsertion = false;
          this.showFail = true;
          this.showSuccess = false;
        }
      );
    }
  }

  configAuths(formData: any) {
    // show the insertion loading
    this.showInsertion = true;
    this.showSuccess = false;
    this.showFail = false;

    const lockArray = [];
    const keyArray = [];

    // get all the locks and keys
    for (let i = 1; i <= this.auths.length; i++) {
      lockArray.push(this.auths[i - 1].toString());
      keyArray.push(formData.value['key' + i]);
    }

    // store keys in plaintext here for usability testing
    this.postConfigFormService.postKeys(this.userID, keyArray, lockArray, this.jwToken).subscribe(
      suc => {
        console.log(suc);
      },
      err => {
        console.log(err );
      }
    );

    // permutate the locks and keys
    const keyPermArray = this.permService.generateLimPerms(keyArray, this.displayLength);

    // hash the lock and key combos
    const hashArray = this.hashService.hashPermsSHA256(keyPermArray);

    // submit the final form
    this.postConfigFormService.postAuthArray(this.userID, lockArray, hashArray, this.jwToken).subscribe(
      suc => {
        this.showInsertion = false;
        this.showSuccess = true;
        this.showFail = false;

        console.log(suc);
        // first initialize the user's storage
        const setInit = JSON.stringify({'init': true});
        localStorage.setItem('currentUser', setInit);

        // then log everything
        const endTime = new Date().getTime();
        const totalCreationTime = endTime - this.loginStartTime; // time to configure
        const avgLength = keyArray.join('').length / keyArray.length; // average length of keys
        const logged: ConfigActivity =  {  userID: Number(this.userID),
                                          totalCreationTime: totalCreationTime,
                                          avgSecretLength: avgLength
                                        } as ConfigActivity;
        this.activityLog.logConfigActivity(logged).subscribe();
        // swap to practice now in the parent component using ngSwitch
      },
      err => {
        this.showInsertion = false;
        this.showSuccess = false;
        this.showFail = true;
        console.log(err );
      }
    );

  }

}
