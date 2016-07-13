//
//  SecondViewController.swift
//  bnwie
//
//  Created by TruongSinh Tran-Nguyen on 7/10/16.
//  Copyright © 2016 TruongSinh Tran-Nguyen. All rights reserved.
//

import UIKit
import FBSDKLoginKit
class LoginVC: UIViewController {
    private var fbLoginSuccess = false
    private let readPermissions = ["public_profile", "email", "user_friends"]
    override func viewDidLoad() {
        super.viewDidLoad()
    }

    override func viewDidAppear(_ animated: Bool) {
        if FBSDKAccessToken.current() != nil || fbLoginSuccess == true {
            performSegue(withIdentifier: "Landing2Main", sender: self)
        }
    }

    @IBAction func doSignInWithFb(sender: FBSDKLoginButton, forEvent event: UIEvent) {
        FBSDKLoginManager()
            .logIn(withReadPermissions: readPermissions, from: self, handler: handleLogin)
    }
    private func handleLogin (_ result: FBSDKLoginManagerLoginResult?, _ err: NSError?) {
            if let err = err {
            return handleLoginError(err)

        }
            guard let result = result else {
            return handleLoginUnexpected()
        }
        if result.isCancelled == true {
            return handleLoginCancel()
        }
        return handleLoginSuccess(result)
    }
    private func handleLoginError(_ err: NSError) {

    }

    private func handleLoginUnexpected() {

    }

    private func handleLoginCancel() {
    }

    private func handleLoginSuccess(_ result: FBSDKLoginManagerLoginResult) {
        self.fbLoginSuccess = true
    }

}
