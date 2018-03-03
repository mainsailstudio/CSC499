import { Component, OnInit } from '@angular/core';
import { RegisterUserService } from '../register/register-user.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-dashboard-nav',
  templateUrl: './dashboard-nav.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardNavComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-dashboard-sidebar',
  templateUrl: './dashboard-sidebar.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardSidebarComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}

@Component({
  selector: 'app-dashboard-main',
  templateUrl: './dashboard-main.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardMainComponent implements OnInit {

  userDataJSON = localStorage.getItem('currentUser');
  userData = JSON.parse(this.userDataJSON);
  email = this.userData['email'];

  constructor() { }

  ngOnInit() {

  }

}

