package customtemplates

import (
	"net/http"

	httperror "github.com/portainer/libhttp/error"
	"github.com/portainer/libhttp/response"
	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/http/security"
	"github.com/portainer/portainer/api/internal/authorization"
)

// @id CustomTemplateList
// @summary List available custom templates
// @description List available custom templates.
// @description **Access policy**: authenticated
// @tags custom_templates
// @security jwt
// @produce json
// @success 200 {array} portainer.CustomTemplate "Success"
// @failure 500 "Server error"
// @router /custom_templates [get]
func (handler *Handler) customTemplateList(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	customTemplates, err := handler.DataStore.CustomTemplate().CustomTemplates()
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve custom templates from the database", err}
	}

	resourceControls, err := handler.DataStore.ResourceControl().ResourceControls()
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve resource controls from the database", err}
	}

	customTemplates = authorization.DecorateCustomTemplates(customTemplates, resourceControls)

	securityContext, err := security.RetrieveRestrictedRequestContext(r)
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve info from request context", err}
	}

	if !securityContext.IsAdmin {
		user, err := handler.DataStore.User().User(securityContext.UserID)
		if err != nil {
			return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve user information from the database", err}
		}

		userTeamIDs := make([]portainer.TeamID, 0)
		for _, membership := range securityContext.UserMemberships {
			userTeamIDs = append(userTeamIDs, membership.TeamID)
		}

		customTemplates = authorization.FilterAuthorizedCustomTemplates(customTemplates, user, userTeamIDs)
	}

	return response.JSON(w, customTemplates)
}
