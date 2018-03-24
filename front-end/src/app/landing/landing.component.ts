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

}
