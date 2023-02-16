import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QuestionComponent } from './question/question.component';
import { ResultsComponent } from './results/results.component';

const routes: Routes = [];

@NgModule({
  imports: [
    RouterModule.forRoot([
      {path: 'question', component: QuestionComponent},
      {path: 'results', component: ResultsComponent},
    ]),
  ],
  exports: [RouterModule
  ]
})
export class AppRoutingModule { }
