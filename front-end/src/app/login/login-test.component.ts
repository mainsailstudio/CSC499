import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { TestURL, APIURL } from '../api/api.constants';
import { Component, OnInit } from '@angular/core';
// import { AuthenticateService } from '../authenticate/authenticate.service';
import { LoginTestService } from './login-test.service';
import { RedirectMessageService } from '../misc/redirect-message.service';
import { ActivityLogService } from '../activity-log/activity-log.service';

import * as shajs from 'sha.js';
import { LogActivity } from '../activity-log/log.interface';

export interface TestUser {
  email: string;
  testLevel: number;
  locks: string;
}

export interface ContTestUser {
  id: string;
  email: string;
  testLevel: number;
  secret: string;
  token: string;
}

@Component({
    selector: 'app-login-test',
    templateUrl: 'login-test.component.html',
    styleUrls: ['./login.component.css']
  })
export class LoginTestComponent implements OnInit {

  initLogin = true;
  testLevel = 0;
  locks = '';
  allowTempPass = false;
  userEmail = '';
  showSuccess = false;
  showFail = false;
  errorMessage = '';

  redirectMessage = this.redirect.message;

  // activity logging variables
  loginStartTime = new Date().getTime();
  failures = 0;
  refreshes = 0;
  secretLength = 0;

  constructor(private loginService: LoginTestService,
              private redirect: RedirectMessageService,
              private activityLog: ActivityLogService,
              private router: Router) { }

  ngOnInit() {
    // const startTime = new Date();
  }

  // function to refresh the user's locks if they are able
  refreshLocks() {
    // first add to the refreshes variable
    this.refreshes++;
    console.log('Refreshes are ' + this.refreshes);
    // since they can't refresh it until postData has already been called once, this should have their email
    this.postData(this.userEmail);
  }

  postData(email: string): void {
   // this.register = undefined;
    email = email.trim();
    this.userEmail = email;
    console.log('Email is ' + this.userEmail);


    const loginUser: TestUser = { email } as TestUser;
    this.loginService.startLoginUser(loginUser).subscribe(
      suc => {
        console.log(suc);
        this.nextFormInput(suc);
      },
      err => {
        console.log(err);
        console.log('Error log here');
      }
    );
  }

  nextFormInput(user: TestUser) {
    this.testLevel = user.testLevel;

    if (user.testLevel === 1) { // test user with passwords
      this.initLogin = false;
    } else if (user.testLevel === 2 || user.testLevel === 3) { // test user with auths
      this.initLogin = false;
      this.locks = user.locks;
    } else {
      this.initLogin = true;
      this.testLevel = 4;
      this.errorMessage = 'Test';
    }
  }

  loginUser(keys: string, tempPass: string) {
    console.log('Keys are ' + keys + ' and tempPass is ' + tempPass);
  }

  // the form that logs in the user regardless if they are using a password or dynauth due to the testLevel variable
  contLogin(contLoginForm: any) {
    const email = this.userEmail;
    const testLevel = this.testLevel;
    let tempPass = contLoginForm.tempPass;
    let auth = contLoginForm.keys;

    if (auth === 'undefined' || auth === undefined) {
      auth = '';
    } else {
      tempPass = '';
      // auth = this.locks + contLoginForm.keys; // testing without locks sent in call, ie session based authentication
    }

    const secretBeforeHash = tempPass + auth;
    this.secretLength = secretBeforeHash.length;
    console.log('Secret length is ' + this.secretLength);
    const secret = shajs('sha256').update(secretBeforeHash).digest('hex'); // due to the if's one or the other should be empty
    const loginUser: ContTestUser = { email, testLevel, secret } as ContTestUser;

    this.loginService.contLoginUser(loginUser).subscribe(
      suc => {
        if (suc) {
          // let the user know that it was successful
          this.showSuccess = true;
          this.showFail = false;

          const endTime = new Date().getTime();
          const loginTime = endTime - this.loginStartTime;
          const userDataJSON = localStorage.getItem('currentUser');
          const userid = Number((JSON.parse(userDataJSON).id));

          // log the login activity
          const logged: LogActivity = { userID: userid,
                                        testLevel: testLevel,
                                        loginTime: loginTime,
                                        failures: this.failures,
                                        refreshes: this.refreshes,
                                        secretLength: this.secretLength
                                       } as LogActivity;
            this.activityLog.logActivity(logged).subscribe();
            console.log('UserID is ' + userid);
            console.log('Called logActivity');

          // 2 second delay before redirecting
          Observable.timer(1200)
          .subscribe(i => {
            this.router.navigate(['/dashboard']);
          });
        } else {
          this.showFail = true;
          this.showSuccess = false;
        }
      },
      err => {
        // add the failure to the failure variable
        this.failures++;
        this.showFail = true;
        this.showSuccess = false;
        console.log('Error is ' + err);
      }
    );
  }

}

