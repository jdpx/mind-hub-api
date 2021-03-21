

function (user, context, callback) {
  console.log("I am hgere");
  context.accessToken['http://foo/bar'] = 'value';
  	context.accessToken['https://mind.jdpx.co.uk/testing'] =  'foodbar';
  
  const rolePrefix = "Organisation"

const testOrgID = "333123ef-342f-42c7-8b25-56ac8b1da008"

const orgMap = {
  "OrganisationTest": testOrgID
}
  console.log("I am hgere1111");

  if (!context.authorization) {
    console.log(`No authorization in event`)

     callback(null, user, context);
    return
  }
  console.log("I am hgere222");

  const orgs = context.authorization.roles.filter(x => x.startsWith(rolePrefix))

  console.log("I am hgere333");
  if (!orgs || orgs.length === 0) {
    console.log(`No orgs in roles`)

    callback(null, user, context);
    return
  }
  console.log("I am hgere444");

  const org = orgs[0]
  const orgID = orgMap[org]

  console.log("I am hgere555");
  if (!orgID) {
    console.log(`No Org ID in Map for ${orgID}`)

    callback(null, user, context);
    return
  }
  console.log("I am hgere666" + orgID);
  
  // context.accessToken['https://mind.jdpx.co.uk/organisation'] =  orgID;  
  
  context.accessToken['https://product-injector-ui.ritech.io/roles'] =  orgID;
  
  console.log("I am hgere777" + JSON.stringify(context));
  
  
  // TODO: implement your rule
  callback(null, user, context);
}