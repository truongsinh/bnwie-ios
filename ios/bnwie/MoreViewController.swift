//
//  SecondViewController.swift
//  bnwie
//
//  Created by TruongSinh Tran-Nguyen on 7/10/16.
//  Copyright © 2016 TruongSinh Tran-Nguyen. All rights reserved.
//

import UIKit
import FBSDKLoginKit

class MoreViewController: UITableViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
    }

    @IBAction func doLogOut() {
            FBSDKLoginManager().logOut()
        performSegue(withIdentifier: "More2Landing", sender: self)
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}
