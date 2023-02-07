import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';

@Component({
  selector: 'question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css'],
})
export class QuestionComponent{
    title = "Questions: ";

    onSubmit(data: string, data2: string) {
      console.log(data, data2);
    }
}
