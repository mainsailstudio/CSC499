import { Component, OnInit } from '@angular/core';
import { RegisterUserService } from '../register/register-user.service';
import { trigger, transition, style, animate, state } from '@angular/animations';
import { AccountInitService } from './account-init.service';

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

  tokenJSON = localStorage.getItem('currentUser');
  testLevel = (JSON.parse(this.tokenJSON).testLevel);
  // loginStateJson = localStorage.getItem('currentUser');
  constructor() { }

  ngOnInit() {
    console.log('Token JSON is ' + this.tokenJSON);
    console.log('Login state is ' + this.testLevel);

    if (this.testLevel === 1) {
        console.log('Password here');
    }
  }

}
