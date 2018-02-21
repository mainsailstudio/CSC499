import { Component, OnInit, Input } from '@angular/core';

@Component({
    selector: 'app-register',
    templateUrl: 'register.component.html',
    styleUrls: ['./register.component.css']
  })
export class RegisterComponent implements OnInit {

  @Input() email: string;

  constructor() { }

  ngOnInit() {
  }

}
