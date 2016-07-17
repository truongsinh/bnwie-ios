//
//  FirstViewController.swift
//  bnwie
//
//  Created by TruongSinh Tran-Nguyen on 7/10/16.
//  Copyright © 2016 TruongSinh Tran-Nguyen. All rights reserved.
//

import UIKit
import Alamofire


let availableService: Array<ServiceStruct>= [
    ServiceStruct(emoji:"💇🏻", value: "hair", text:NSLocalizedString("Hair care", comment: "")),
    ServiceStruct(emoji:"💄", value: "face", text:NSLocalizedString("Facial care", comment: "")),
    ServiceStruct(emoji:"💅🏼", value: "nail", text:NSLocalizedString("Manicure & Pedicure", comment: "")),
    ServiceStruct(emoji:"👙", value: "body", text:NSLocalizedString("Body care", comment: "")),
]

let numAvailableService = availableService.count

class MyProfileViewController: UITableViewController {
    
    @IBOutlet weak var profileImage: UIImageView!
    var providerSectionIndex = 0
    var consumerSectionIndex = 0
    override func viewDidLoad() {
        super.viewDidLoad()
        
        // Do any additional setup after loading the view, typically from a nib.
    }
    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        let numbOfSection = super.numberOfSectionsInTableView(tableView)
        consumerSectionIndex = numbOfSection - 1
        providerSectionIndex = numbOfSection - 2
        return numbOfSection
    }
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        switch section {
        case consumerSectionIndex, providerSectionIndex:
            return numAvailableService
        default:
            return super.tableView(tableView, numberOfRowsInSection: section)
        }
    }
    
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        switch indexPath.section {
        case consumerSectionIndex, providerSectionIndex:
            let cell = ToggleTableViewCell(
                service: availableService[indexPath.row],
                on: true
            )
            return cell
        default:
            return super.tableView(tableView, cellForRowAtIndexPath: indexPath)
        }
        
    }
    
    
    internal func preapreImage(url: String) {
        Alamofire
            .request(
                .GET,
                url
            )
            .responseData{ response in
                guard let imagedData = response.data else{
                    print("return no data")
                    return
                }
                guard let i = UIImage(data: imagedData) else{
                    print("return no image")
                return
                }
                print("set image")
                        self.profileImage?.image = i
        }
        
    }
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
        
    
}
