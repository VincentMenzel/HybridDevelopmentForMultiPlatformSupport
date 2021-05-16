import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup} from '@angular/forms';
import {AuthService} from '../../shared/service/auth.service';
import {NavController} from '@ionic/angular';
import {ToastService} from "../../shared/service/toast.service";

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.page.html',
  styleUrls: ['./sign-in.page.scss'],
})
export class SignInPage implements OnInit {
  public loginForm: FormGroup;

  constructor(formBuilder: FormBuilder, private authService: AuthService, private navController: NavController, private toastService: ToastService) {
    this.loginForm = formBuilder.group({
      username: [''],
      password: ['']
    });
  }

  ngOnInit() {
  }

  onSignIn($event: MouseEvent) {

    this.authService.authenticate({
        username: this.loginForm.value.username,
        password: this.loginForm.value.password
      })
      .subscribe(() => {

        this.navController.navigateForward('signed-in');
        this.toastService.presentToast('Login Erfolgreich', 1000);

      }, (err) => {
        if (err.status === 401) {

          this.toastService.presentToast('Login Ungültig', 1000);

        } else {

          this.toastService.presentToast('Etwas ist schief gelaufen!', 1000);

        }
      });
  }
}
