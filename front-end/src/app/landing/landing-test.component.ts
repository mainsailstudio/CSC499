import { Component, OnInit, ViewChild, AfterViewInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Observable} from 'rxjs/Observable';
// import { RegisterStartComponent, RegisterContinueComponent } from '../register/register.component';
// import { LoginComponent, LoginSuccessComponent } from '../login/login.component';

@Component({
  selector: 'app-landing-test',
  templateUrl: './landing-test.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingTestComponent implements OnInit {

  constructor() { }

  ngOnInit() {
    // this.getUsers();
  }

}
