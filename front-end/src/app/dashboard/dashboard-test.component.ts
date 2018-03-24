import { Component, OnInit } from '@angular/core';
import { RegisterUserService } from '../register/register-user.service';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { Observable } from 'rxjs/Observable';
import * as zxcvbn from 'zxcvbn';

import { DashboardComponent } from './dashboard.component';

import { PermutateService } from '../hash/perm.service';
import { HashSha256Service } from '../hash/hash-sha256.service';
import { CombinePermsService } from '../hash/combine.service';
import { InitAccountService } from './init-account.service';
import { ActivityLogService } from '../activity-log/activity-log.service';

import { ConfigActivity } from '../activity-log/log.interface';

import * as shajs from 'sha.js';
import { UserConstantsService } from './user-constants/user-constants.service';
import { RegisterTestUser } from '../register/register-test.component';
import { RegisterTestService } from '../register/register-test.service';
import { GetTokenService } from '../misc/get-token.service';
import { Router } from '@angular/router';

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

  public account = {
    password: null
  };
  public barLabel = 'Password strength:';

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

  // this is the success variable that shows after the user account is configured
  showInsertion = false;
  showSuccess = false;
  showFail = false;
  showLengthError = false;
  showPasswordError = false;
  showAuthLengthError = false;

  // activity logging variables
  loginStartTime = new Date().getTime();

  constructor(
    private permService: PermutateService,
    private hashService: HashSha256Service,
    private combineService: CombinePermsService,
    private postConfigFormService: InitAccountService,
    private activityLog: ActivityLogService,
    private registerService: RegisterTestService,
    private getMeThatThereTokenPleaseSir: GetTokenService, // I hate myself sometimes
    private router: Router
  ) { }

  validateRegistration(pass: string, confirmPass: string): boolean {
    if (pass === confirmPass) {
      return true;
    } else {
      return false;
    }
  }

  ngOnInit() {
    // get token data
    this.userDataJSON = localStorage.getItem('currentUser');
    // parse that data
    this.userData = JSON.parse(this.userDataJSON);
    this.init = this.userData['init'];
    if (this.testLevel === 1) {
    } else if (this.testLevel === 2) {
        this.auths = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    } else if (this.testLevel === 3) {
        this.auths = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    }
  }

  configPass(formData: any) {
    this.showPasswordError = false;
    // verify the pass and confirm pass are the same again
    const password = formData.value['tempPass'];
    const confirmPass = formData.value['confirmPass'];
    const zxcvbnResult = zxcvbn(password);
    console.log('Zxcvbn result is ' + JSON.parse(zxcvbnResult.score));

    if (password !== confirmPass || JSON.parse(zxcvbnResult.score) < 3) {
      this.showPasswordError = true;
      return;
    } else {
      this.showPasswordError = false;
      this.showInsertion = true;
      const hashedPass = shajs('sha256').update(password).digest('hex');
      this.postConfigFormService.postPassword(this.userID, hashedPass, this.jwToken).subscribe(
        suc => {
          this.showInsertion = false;
          this.showSuccess = true;
          this.showFail = false;

          // first initialize the user's storage
          const setInit = JSON.stringify({'init': true});
          localStorage.setItem('currentUser', setInit);

          // then log everything
          const endTime = new Date().getTime();
          const totalCreationTime = endTime - this.loginStartTime; // time to configure
          const logged: ConfigActivity =  {  userID: Number(this.userID),
                                            totalCreationTime: totalCreationTime,
                                            avgSecretLength: password.length
                                          } as ConfigActivity;
          this.activityLog.logConfigActivity(logged).subscribe();

          // get a new token to prevent errors
          const testUser: RegisterTestUser = { email: this.email } as RegisterTestUser;
          this.registerService.registerTest(testUser).subscribe(
            success => {
              this.showSuccess = true;
              this.showFail = false;

               // get a new token to prevent errors
              const newTokenUser: RegisterTestUser = { email: this.email } as RegisterTestUser;
              this.getMeThatThereTokenPleaseSir.getFussFreeToken(newTokenUser).subscribe(
                successAgain => {
                  this.showSuccess = true;
                  this.showFail = false;
                },
                error => {
                  this.showSuccess = false;
                  this.showFail = true;
                    console.log(error);
                }
              );

              // 1.2 second delay before redirecting to practice
              Observable.timer(1200)
              .subscribe(i => {
                this.router.navigate(['/practice']);
              });
            },
            error => {
              this.showSuccess = false;
              this.showFail = true;
                console.log(error);
            }
          );
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
    this.showLengthError = false;
    this.showAuthLengthError = false;

    // start by validating that the keys are long enough and that there are no duplicates
    const lockArray = [];
    const keyArray = [];

    // get all the locks and keys
    for (let i = 1; i <= this.auths.length; i++) {
      const key = formData.value['key' + i];

      // make sure they are least 4 characters
      if (key.length < 3) {
        this.showInsertion = false;
        this.showLengthError = true;
        return;
      }

      lockArray.push(this.auths[i - 1].toString());
      keyArray.push(key);
    }

    // make sure there are no duplicates in the key array
    const valuesSoFar = Object.create(null);
    for (let i = 0; i < keyArray.length; ++i) {
        const value = keyArray[i];
        if (value in valuesSoFar) {
          // there are duplicates
          this.showInsertion = false;
          this.showAuthLengthError = true;
          return;
        }
        valuesSoFar[value] = true;
    }
    // there are NO duplicates
    this.showAuthLengthError = false;

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

        // get a new token to prevent errors
        const testUser: RegisterTestUser = { email: this.email } as RegisterTestUser;
        this.getMeThatThereTokenPleaseSir.getFussFreeToken(testUser).subscribe(
          success => {
            this.showSuccess = true;
            this.showFail = false;
          },
          error => {
            this.showSuccess = false;
            this.showFail = true;
              console.log(error);
          }
        );

        // 1.2 second delay before redirecting to practice
        Observable.timer(1200)
        .subscribe(i => {
           this.router.navigate(['/practice']);
        });
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
