function Profile_RenderLinks( links ) {
    const icon_table = [
        ["twitter", "http://pluspng.com/img-png/twitter-png-logo-logo-twitter-in-png-2500.png"],
        ["github", "https://pngimg.com/uploads/github/github_PNG84.png"],
        ["instagram", "https://freepngimg.com/download/logo/69823-instagram-icons-symbol-youtube-computer-logo.png"],
    ];


    let panda_links = document.getElementsByClassName("panda-profile-links")[0];

    for (let i = 0; i < links.length; i++) {
        let name = links[i].url;
        let url = links[i].url;

        let panda_link = document.createElement("div");
        panda_link.className = "panda-link";

        let icon_source = "http://clipart-library.com/img/1930942.png";

        for (let j = 0; j < icon_table.length; j++) {
            let service_name = icon_table[j][0];
            let source = icon_table[j][1];

            if (url.contains(service_name)) {
                icon_source = source;
                break;
            }
        }

        var panda_link_icon = document.createElement("img");
        var panda_link_title = document.createElement("p");

        panda_link_icon.setAttribute("src", icon_source);
        panda_link_icon.height = 30;

        panda_link_title.innerText = name;

        panda_link.appendChild(panda_link_icon);
        panda_link.appendChild(panda_link_title);


        panda_links.appendChild(panda_link);
    }

}

function Profile_RenderPublicUser( public_user ) {

    console.log(public_user);

    let panda_img = document.getElementById("panda-profile-image");
    let panda_name = document.getElementById("panda-profile-name");
    let panda_handle = document.getElementById("panda-profile-handle");
    let panda_bio = document.getElementById("panda-profile-bio");
    
    panda_img.setAttribute("src", public_user.profile_photo_b64);
    panda_name.innerText = public_user.name;
    panda_handle.innerText = `@${public_user.username}`;
    panda_bio.innerText = public_user.bio;

    
    Profile_RenderLinks( public_user.links );

    let primary_color_hex = public_user.primary_color.toString(16);
    let primary_color_hex0 = "0".repeat(6 - primary_color_hex.length) + primary_color_hex;

    let secondary_color_hex = public_user.secondary_color.toString(16);
    let secondary_color_hex0 = "0".repeat(6 - secondary_color_hex.length) + secondary_color_hex;
    
    document.documentElement.style.setProperty("--color0", `#${primary_color_hex0}`);
    document.documentElement.style.setProperty("--color1", `#${secondary_color_hex0}`);

}

function Dashboard_RenderLinks( links ) {
    const icon_table = [
        ["twitter", "http://pluspng.com/img-png/twitter-png-logo-logo-twitter-in-png-2500.png"],
        ["github", "https://pngimg.com/uploads/github/github_PNG84.png"],
        ["instagram", "https://freepngimg.com/download/logo/69823-instagram-icons-symbol-youtube-computer-logo.png"],
    ];


    let panda_links = document.getElementsByClassName("panda-dashboard-profile-links")[0];

    for (let i = 0; i < links.length; i++) {
        let name = links[i].name;
        let url = links[i].url;

        let panda_link = document.createElement("div");
        panda_link.className = "panda-link";

        let icon_source = "http://clipart-library.com/img/1930942.png";

        for (let j = 0; j < icon_table.length; j++) {
            let service_name = icon_table[j][0];
            let source = icon_table[j][1];

            if (url.contains(service_name)) {
                icon_source = source;
                break;
            }
        }

        let panda_link_icon = document.createElement("img");


        let panda_dashboard_name = document.createElement("div");

        let dashboard_name_p = document.createElement("p");
        dashboard_name_p.innerText = name;

        let dashboard_name_url = document.createElement("p");
        dashboard_name_url.className = "panda-url-view-text";
        dashboard_name_url.innerText = url;

        panda_dashboard_name.appendChild(dashboard_name_p);
        panda_dashboard_name.appendChild(dashboard_name_url);


        let panda_link_terminate = document.createElement("div");
        panda_link_terminate.className = "panda-link-terminate";

        let terminate_img = document.createElement("img");
        terminate_img.setAttribute("src", "https://www.shareicon.net/data/2016/01/03/697421_button_512x512.png");
        terminate_img.height = 20;
        panda_link_terminate.appendChild(terminate_img);


        panda_link.appendChild(panda_link_icon);
        panda_link.appendChild(panda_dashboard_name);
        panda_link.appendChild(panda_link_terminate);


        panda_links.appendChild(panda_link);
    }
}

function Dashboard_RenderAdminsterUser( admin_user ) {
    let panda_img = document.getElementById("panda-profile-image");
    let panda_login_msg = document.getElementById("panda-login-message");
    let panda_name = document.getElementById("panda-account-name");
    let panda_email = document.getElementById("panda-account-email");
    let panda_handle = document.getElementById("panda-account-username");
    let panda_bio = document.getElementById("panda-account-bio");
    let panda_begin = document.getElementById("panda-account-begin");
    let panda_end = document.getElementById("panda-account-end");
    

    let primary_color_hex = admin_user.primary_color.toString(16);
    let primary_color_hex0 = "0".repeat(6 - primary_color_hex.length) + primary_color_hex;

    let secondary_color_hex = admin_user.secondary_color.toString(16);
    let secondary_color_hex0 = "0".repeat(6 - secondary_color_hex.length) + secondary_color_hex;

    console.log(primary_color_hex0, secondary_color_hex0);

    document.documentElement.style.setProperty("--color0", `#${primary_color_hex0}`);
    document.documentElement.style.setProperty("--color1", `#${secondary_color_hex0}`);

    panda_img.setAttribute("src", admin_user.profile_photo_b64);
    panda_login_msg.innerText = `Welcome, ${admin_user.name}.`;

    panda_name.value = admin_user.name;
    panda_handle.value = admin_user.username;
    panda_bio.value = admin_user.bio;
    panda_email.value = admin_user.email;
    panda_begin.value = `#${primary_color_hex0}`;
    panda_end.value = `#${secondary_color_hex0}`;

    Dashboard_RenderLinks( admin_user.links );
} 