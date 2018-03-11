import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';


import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { HttpErrorHandler, HandleError } from '../http-error-handler.service';

import { StartLoginUser, ContLoginUser } from './login.component';
import { TestURL, APIURL } from '../api/api.constants';
import { Component, OnInit } from '@angular/core';
// import { AuthenticateService } from '../authenticate/authenticate.service';
import { LoginTestService } from './login-test.service';

export interface StartLoginUser {
  email: string;
  loginState: string;
  locks: string;
}

export interface ContLoginUser {
  id: string;
  email: string;
  loginState: string;
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
  loginState = '0';
  locks = '';
  allowTempPass = false;
  userEmail = '';
  showSuccess = false;
  showFail = false;
  constructor(private loginService: LoginTestService) { }

  ngOnInit() {
    // this.getHeroes();
  }

  postData(email: string): void {
   // this.register = undefined;
    email = email.trim();
    this.userEmail = email;
    console.log('Email is ' + this.userEmail);


    const loginUser: StartLoginUser = { email } as StartLoginUser;
    console.log('loginUser to be posted is ' + loginUser);
    this.loginService.startLoginUser(loginUser).subscribe(
      suc => {
        console.log(suc);
        this.nextFormInput(suc);
      },
      err => {
        console.log(err );
      }
    );
  }

  nextFormInput(user: StartLoginUser) {
    this.loginState = user.loginState;

    if (user.loginState === '1') { // first time user login
      this.initLogin = false;
    } else if (user.loginState === '2') { // user hasn't initialized locks yet
      this.initLogin = false;
      this.locks = user.locks;
    } else if (user.loginState === '3') { // user is all set
      this.initLogin = false;
      this.locks = user.locks;
    } else {
      this.initLogin = true; // not necessary but for safety
      this.loginState = 'Unknown';
    }
  }

  loginUser(keys: string, tempPass: string) {
    console.log('Keys are ' + keys + ' and tempPass is ' + tempPass);
  }

  contLogin(contLoginForm: any) {
    const email = this.userEmail;
    let loginState = this.loginState;
    let tempPass = contLoginForm.tempPass;
    let auth = contLoginForm.keys;

    console.log('Email is ' + email);
    console.log('loginState is ' + loginState);
    console.log('tempPass is ' + tempPass);
    console.log('auth is ' + auth);
    // for documentation purposes:
    // this uses the same loginState variable to determine if the user is logging in via a temp pass or their lock-key combo.
    // Since there is only those 2 options, and loginState of 2 initially determines that the user could use either,
    // these if statements set the loginState to be either 1 (password) or 3 (dynauth)
    if (auth === 'undefined' || auth === undefined) {
      auth = '';
      loginState = '1';
    } else {
      tempPass = '';
      loginState = '3';
      auth = this.locks + contLoginForm.keys;
    }

    const secret = tempPass + auth; // due to the if's one or the other should be empty
    console.log('Secret is ' + secret + ' and loginstate is ' + loginState);
    const loginUser: ContLoginUser = { email, loginState, secret } as ContLoginUser;
    console.log('loginUser to be posted is ' + email + tempPass + auth);
    this.loginService.contLoginUser(loginUser).subscribe(
      suc => {
        if (suc) {
// new Headers({ 'Authorization': 'Bearer ' + this.authenticationService.token });
          this.showSuccess = true;
          this.showFail = false;
          console.log('Success');
        } else {
          this.showFail = true;
          this.showSuccess = false;
          console.log('Failure');
        }
      },
      err => {
        this.showFail = true;
        this.showSuccess = false;
        console.log('Error is ' + err);
      }
    );
    // console.log('contLoginForm is ' + contLoginForm.tempPass);
  }


  toggleTempPass() {
    if (this.allowTempPass === false) {
      this.allowTempPass = true;
    } else {
      this.allowTempPass = false;
    }
  }

  tryPing() {

  }

}

