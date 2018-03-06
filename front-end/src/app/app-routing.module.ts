import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RegisterStartComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { LogoutComponent } from './login/logout.component';
import { LandingComponent } from './landing/landing.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { DashboardComponent } from './dashboard/dashboard.component';

// new imports for authentication
import { LoginNewComponent } from './login/login-new.component';
import { AuthGuard } from './_auth-guard/auth.guard';

const routes: Routes = [
  { path: '',  component: LandingComponent },
  { path: 'register', component: RegisterStartComponent },
  { path: 'login', component: LoginNewComponent },
  { path: 'login-real', component: LoginComponent },
  { path: 'logout', component: LogoutComponent },
  { path: 'dashboard', component: DashboardComponent, canActivate: [AuthGuard] },
  { path: '**', component: NotFoundComponent } // make sure this is always at the bottom so it doesn't superscede legitimate routes
];

@NgModule({
  imports: [
     RouterModule.forRoot(routes)
    ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }

