import {Injectable} from '@angular/core';
import {BehaviorSubject, Observable} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {environment} from '../../../environments/environment';
import {tap} from 'rxjs/operators';
import {IUser} from "../interface/user";
import {ToastController} from "@ionic/angular";
import {ToastService} from "./toast.service";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  onUserChange = new BehaviorSubject<null | IUser>(null);

  constructor(
    private http: HttpClient,
    private toastService: ToastService
  ) {
  }

  public authenticate({username, password}: { username: string; password: string }) {
    console.log(`${environment.apiUrl}/signIn`);
    return this.http.post<IUser>(`${environment.apiUrl}/signIn`, {username, password}, {
        headers: new HttpHeaders({
          "Content-Type": "text/plain"
        })
      })
      .pipe(
        tap(user => {
          this.onUserChange.next(user);
        }, (err) => {
          console.error(err);
        })
      );
  }
}
