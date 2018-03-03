import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Observable} from 'rxjs/Observable';
import {TestService} from './test.service';
// import { RegisterStartComponent, RegisterContinueComponent } from '../register/register.component';
// import { LoginComponent, LoginSuccessComponent } from '../login/login.component';

@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingComponent implements OnInit {

  loginState = 1; // initialize login state
  registerState = 1; // initialize register state

  constructor() { }

  ngOnInit() {
    // this.getUsers();
  }

  changeLoginState(state: number) {
    console.log('changing view to ' + event);
    this.loginState = state;
  }

  changeRegisterState(state: number) {
    console.log('changing view to ' + event);
    this.registerState = state;
  }

  // register(): void {
  //   if (this.displayRegisterStart === true) {

  //     this.displayRegisterStart = false;
  //     this.displayRegisterContinue = true;
  //   } else if (this.displayRegisterContinue === true) {
  //       this.displayRegisterContinue = false;
  //       this.displayRegisterFinal = true;

  //   }
  // }

  // getUsers() {
  //   this._testService.getUsers().subscribe(
  //     // the first argument is a function which runs on success
  //     data => { this.users = data; },
  //     // the second argument is a function which runs on error
  //     err => console.error(err),
  //     // the third argument is a function which runs on completion
  //     () => console.log('done loading users')
  //   );
  // }

}
