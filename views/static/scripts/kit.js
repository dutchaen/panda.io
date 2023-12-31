function isValidColor( color_text )  {

    let allowed_chars = "0123456789abcdef";
    let x = color_text.toLowerCase().replaceAll("#", '');
    if (x.length !== 3 && x.length !== 6) {
        return false;
    }

    for (let i = 0; i < x.length; i++) {
        if (!allowed_chars.includes(x.charAt(i))) {
            return false;
        }
    }

    return true;
}

function isNameValid( name ) {
    return name.length > 0 && name.length <= 25;
} 

function isEmailValid( email ) {
    const expr = new RegExp("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])");
    return expr.test(email);
}

function isUsernameValid( username ) {
    return username.length > 0 && username.length <= 25;
}


function convertHexColorToInt( color_text ) {

    if (!isValidColor( color_text )) {
        return undefined;
    }

    let hex_string = "0x";
    let color_int = undefined;
    let x = color_text.toLowerCase().replaceAll("#", '');

    switch (x.length) {
        case 3:
            for (let i = 0; i < 3; i++) {
                hex_string += x.charAt(i).repeat(2);
            }
            color_int = parseInt(hex_string, 16);
            break;
        case 6:
            color_int = parseInt(x, 16);
            break;
        default:
            break;

    }

    return color_int;
} 
