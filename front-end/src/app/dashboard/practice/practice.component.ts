import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../../http-error-handler.service';

import { TestURL, APIURL } from '../../api/api.constants';
import { Component, OnInit } from '@angular/core';
// import { AuthenticateService } from '../authenticate/authenticate.service';
import { LoginTestService } from '../../login/login-test.service';
import { RedirectMessageService } from '../../misc/redirect-message.service';
import { ActivityLogService } from '../../activity-log/activity-log.service';

import * as shajs from 'sha.js';
import { LoginActivity } from '../../activity-log/log.interface';
import { TestUser, ContTestUser } from '../../login/login-test.component';
import { PracticeService } from './practice.service';
import { UserConstantsService } from '../user-constants/user-constants.service';

@Component({
    selector: 'app-dashboard-main-practice',
    templateUrl: 'practice.component.html',
    styleUrls: ['./practice.component.css']
  })
export class PracticeComponent implements OnInit {

  // get user's keys into this
  keys = [];
  password = '';

  // will be filled by postData
  locks = [];
  locksString = '';

  // variable for authInput (for clearing on refresh)
  authInput = '';

  // variables to programatically display
  showLoading = false;
  showSuccess = false;
  showFail = false;

  // activity logging variables
  loginStartTime = new Date().getTime();
  failures = 0;
  refreshes = 0;
  secretLength = 0;

  constructor(private loginService: LoginTestService,
              private activityLog: ActivityLogService,
              private practiceService: PracticeService,
              public userConstants: UserConstantsService
              ) { }

  ngOnInit() {
    if (this.userConstants.TestLevel === 1) {
      this.practiceService.getTestUserDisplayPassword(this.userConstants.ID).subscribe(
        suc => {
          this.password = suc;
        },
        err => {
        }
      );
    }

    if (this.userConstants.TestLevel === 2 || this.userConstants.TestLevel === 3) {
      this.practiceService.getTestUserKeys(this.userConstants.ID).subscribe(
        suc => {
          this.keys = Array.from(suc); // what in the hell is going on with this, idk internet code works
        },
        err => {
        }
      );
    }
    this.postData(this.userConstants.Email);
  }

  // function to refresh the user's locks if they are able
  refreshLocks() {
    // first add to the refreshes variable
    this.refreshes++;
    this.authInput = '';
    this.showLoading = false;
    this.showSuccess = false;
    this.showFail = false;
    // since they can't refresh it until postData has already been called once, this should have their email
    this.postData(this.userConstants.Email);
  }

  postData(email: string): void {
   // this.register = undefined;
    email = email.trim();
    this.userConstants.Email = email;

    const loginUser: TestUser = { email } as TestUser;
    this.loginService.startLoginUser(loginUser).subscribe(
      suc => {
        this.nextFormInput(suc);
      },
      err => {
      }
    );
  }

  nextFormInput(user: TestUser) {
    this.userConstants.TestLevel = user.testLevel;

    if (user.testLevel === 1) { // test user with passwords
      this.userConstants.Init = false;
    } else if (user.testLevel === 2 || user.testLevel === 3) { // test user with auths
      this.userConstants.Init = false;
      this.locks = user.locks;
      this.locksString = this.locks.join(' - ');
    } else {
      this.userConstants.Init = true;
      this.userConstants.TestLevel = 4;
    }
  }

  loginUser(keys: string, tempPass: string) {
  }

  // the form that logs in the user regardless if they are using a password or dynauth due to the testLevel variable
  contLogin(contLoginForm: any) {

    // initialize notifications
    this.showSuccess = false;
    this.showFail = false;
    this.showLoading = true;

    const email = this.userConstants.Email;
    const testLevel = this.userConstants.TestLevel;
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
    const secret = shajs('sha256').update(secretBeforeHash).digest('hex'); // due to the if's one or the other should be empty
    const loginUser: ContTestUser = { email, testLevel, secret } as ContTestUser;
    this.loginService.contLoginUser(loginUser).subscribe(
      suc => {
        if (suc) {
          // let the user know that it was successful
          this.showLoading = false;
          this.showSuccess = true;
          this.showFail = false;

          const endTime = new Date().getTime();
          const loginTime = endTime - this.loginStartTime;
          const userDataJSON = localStorage.getItem('currentUser');
          const userid = Number((JSON.parse(userDataJSON).id));

          // log the login activity
          const logged: LoginActivity = { userID: userid,
                                        testLevel: testLevel,
                                        loginTime: loginTime,
                                        failures: this.failures,
                                        refreshes: this.refreshes,
                                        secretLength: this.secretLength
                                       } as LoginActivity;
          this.activityLog.logPracticeActivity(logged).subscribe();

          // this is where the practice logging differs from the actual logging, all variables are reset
          this.loginStartTime = new Date().getTime();
          this.failures = 0;
          this.refreshes = 0;
          this.secretLength = 0;

        } else {
          this.showLoading = false;
          this.showFail = true;
          this.showSuccess = false;
        }
      },
      err => {
        // add the failure to the failure variable
        this.failures++;
        this.showLoading = false;
        this.showFail = true;
        this.showSuccess = false;
      }
    );
  }

}

