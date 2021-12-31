import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { FormlyFieldConfig } from '@ngx-formly/core';
import { ConfigService } from 'src/app/services/config.service';
import { ToastrService } from 'ngx-toastr';
import { ErrorMsgService } from 'src/app/services/error-msg.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  form = new FormGroup({});
  fields: FormlyFieldConfig[] = [];
  model:SignupModel = { email: '', password:'' , repassword: ''};

  constructor(
    private http: HttpClient,
    private router: Router,
    private toastr: ToastrService,
    private configService: ConfigService,
    private errMsgService: ErrorMsgService,
    ) {
      this.initFormly();
     }

  ngOnInit(): void {
  }

  submit() {
    if (!this.form.valid) {
      return
    }

    if (this.model.password !== this.model.repassword) {
      this.toastr.error('Please enter the same password', 'Error!');
      return;
    }
    
    const postData = {
      user : this.model.email,
      password: this.model.password,
    }
    this.http.post(`${this.configService.baseUrl()}/api/v1/signup`, postData)
      .subscribe({
        next: (v) => {
          alert(JSON.stringify(v));
          this.router.navigateByUrl('/signin');
        },
        error: (e) => {
          this.toastr.error(this.errMsgService.get(e.error), 'Error!');
          console.error(e)
        },
        complete: () => {}
    });
  }

  private initFormly(){
    this.fields = [
      {
        key: 'email',
        type: 'input',
        templateOptions: {
          label: 'Email Address',
          placeholder: 'Enter Email',
          required: true,
        }
      },
      {
        key: 'password',
        type: 'input',
        templateOptions: {
          label: 'Password',
          placeholder: 'Enter Password',
          required: true,
        }
      },
      {
        key: 'repassword',
        type: 'input',
        templateOptions: {
          label: 'Re-enter Password',
          placeholder: 'Re-enter Password',
          required: true,
        }
      },
    ];
  }
}

interface SignupModel {
  email: string;
  password: string;
  repassword: string;
}
