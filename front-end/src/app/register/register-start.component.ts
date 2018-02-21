import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-register-start',
  templateUrl: './register-start.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterStartComponent implements OnInit {

  @Output() messageEvent = new EventEmitter<string>();

  constructor() { }

  ngOnInit() {
  }

  sendEmail() {
    this.messageEvent.emit('test');
  }

}
