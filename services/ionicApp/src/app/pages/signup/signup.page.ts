import {Component, OnInit} from '@angular/core';
import {Form, FormBuilder, FormGroup} from '@angular/forms';
import {ToastService} from '../../shared/service/toast.service';
import {UserService} from '../../shared/service/user.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.page.html',
  styleUrls: ['./signup.page.scss'],
})
export class SignupPage implements OnInit {
  private registerForm: FormGroup;

  constructor(formBuilder: FormBuilder, private toastService: ToastService, private userService: UserService) {
    this.registerForm = formBuilder.group({
      username: ['',],
      password: ['',],
      passwordRepeat: ['',],
    });
  }

  ngOnInit() {
  }

  onSignup($event: MouseEvent) {
    const {username, password, passwordRepeat} = this.registerForm.value;

    console.log({username, password, passwordRepeat}, this.registerForm.value);
    if (username?.length === 0) {
      return this.toastService.presentToast('Username benötigt', 1000);
    }

    if (password?.length === 0) {
      return this.toastService.presentToast('Passwor benötigt', 1000);
    }

    if (password !== passwordRepeat) {
      return this.toastService.presentToast('Passwörter müssen übereinstimmen', 1000);
    }


    this.userService.postUser({username, password})
      .subscribe(() => {
      this.toastService.presentToast('Registrierung erfolgreich', 1000);
    }, (err) => {
        console.error(err);
        this.toastService.presentToast('Etwas ist schief gelaufen', 1000);
      });
  }

  private onSignupSuccess() {
  }

  private onSignupFailure(err) {
  }
}
