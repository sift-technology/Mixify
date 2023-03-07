import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QuestionComponent } from './question/question.component';
import { ResultsComponent } from './results/results.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [];

@NgModule({
  imports: [
    RouterModule.forRoot([
      {path: 'login', component: LoginComponent},
      {path: '', component: QuestionComponent},
      {path: 'results', component: ResultsComponent},
    ]),
  ],
  exports: [RouterModule
  ]
})
export class AppRoutingModule { }
