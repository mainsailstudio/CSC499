// package imports
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

// app component imports
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { LoginStartComponent } from './login/login-start.component';
import { RegisterComponent } from './register/register.component';
import { RegisterStartComponent } from './register/register-start.component';
import { LandingComponent } from './landing/landing.component';
import { NotFoundComponent } from './not-found/not-found.component';

// module imports
import { AppRoutingModule } from './app-routing.module';

// service imports
import { LoginService } from './login/login.service';
import { TestService } from './landing/test.service';
import { RegisterStartService } from './register/register-start.service';
import { HttpErrorHandler } from './http-error-handler.service';
import { MessageService } from './message.service';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    LoginStartComponent,
    RegisterComponent,
    RegisterStartComponent,
    LandingComponent,
    NotFoundComponent
   ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    NgbModule.forRoot()
  ],
  providers: [
    LoginService,
    TestService,
    RegisterStartService,
    HttpErrorHandler,
    MessageService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
