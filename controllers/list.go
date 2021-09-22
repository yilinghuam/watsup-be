package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"watsup.com/models"
)

type ViewIndividualGroupbuyConfig struct {
	GroupbuyInfo *models.Groupbuy    `json:"GroupbuyInfo"`
	HostItemInfo []models.ItemConfig `json:"HostItemInfo"`
}
type AddIndividualOrderConfig struct {
	Address string                 `json:"Address"`
	Order   []models.AddItemConfig `json:"Order"`
}
type ViewOrderConfig struct {
	HostInfo []models.ItemConfig `json:"HostInfo"`
	UserInfo []models.ItemConfig `json:"UserInfo"`
}
type StatusConfig struct {
	Status string `json:"Status"`
}

func DeleteOrder(c echo.Context) error {
	order_id, _ := strconv.Atoi(c.Param("id"))

	orderErr := models.DeleteOrdersByOrderId(order_id)
	if orderErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "delete order issue")
	}
	itemErr := models.DeleteItemsByOrderId(order_id)
	if itemErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "delete item issue")
	}

	return c.JSON(http.StatusOK, "deleted")
}

func ChangeOrderStatus(c echo.Context) error {
	order_id, _ := strconv.Atoi(c.Param("id"))
	var status StatusConfig
	if err := c.Bind(&status); err != nil {
		return err
	}
	fmt.Println(1212)
	fmt.Println(status)
	statuserr := models.ChangeOrderStatusByOrderID(status.Status, order_id)
	if statuserr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cannot change order status")
	}
	return c.JSON(http.StatusOK, "Order updated!")
}

func GetOrderStatus(c echo.Context) error {
	groupbuy_id, _ := strconv.Atoi(c.Param("id"))

	order, err := models.GetStatusByGroupbuy(groupbuy_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cannot get order status")
	}
	return c.JSON(http.StatusOK, order)
}

func ChangeGroupbuyStatus(c echo.Context) error {
	groupbuy_id, _ := strconv.Atoi(c.Param("id"))
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user := claims["name"].(string)
	var status StatusConfig
	if err := c.Bind(&status); err != nil {
		return err
	}
	fmt.Println(1212)
	fmt.Println(status)
	statuserr := models.ChangeStatusByGroupbuyID(user, status.Status, groupbuy_id)
	if statuserr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cannot change status")
	}
	return c.JSON(http.StatusOK, "Groupbuy updated!")
}

// groupbuy page
func DeleteGroupbuy(c echo.Context) error {
	groupbuy_id, _ := strconv.Atoi(c.Param("id"))
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user := claims["name"].(string)
	fmt.Println(user)

	// check if groupbuy belongs to user
	ownerErr := models.CheckGroupbuyOwner(user, groupbuy_id)

	if ownerErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "not owner")
	}
	orderErr := models.DeleteOrdersByGroupbuy(groupbuy_id)
	if orderErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "delete order issue")
	}
	itemErr := models.DeleteItemsByGroupbuy(groupbuy_id)
	if itemErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "delete item issue")
	}
	groupbuyErr := models.DeleteGroupbuyInfo(user, groupbuy_id)
	// if owner then delete all

	// if not owner, then stop
	if groupbuyErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "not owner")
	}

	return c.JSON(http.StatusOK, "deleted")
}
func ViewOrders(c echo.Context) error {
	groupbuy_id, _ := strconv.Atoi(c.Param("id"))
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user := claims["name"].(string)
	fmt.Println(user)
	// check if groupbuy belongs to user

	ownerErr := models.CheckGroupbuyOwner(user, groupbuy_id)

	// if not owner, get only own info
	if ownerErr != nil {
		userInfo, err := models.GetUserItems(groupbuy_id)
		if err != nil {
			return echo.NewHTTPError(http.StatusOK, userInfo)
		}
	}
	// if owner then get others
	AllInfo, err := models.GetUserItems(groupbuy_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Get host items issue")
	}

	return c.JSON(http.StatusOK, AllInfo)
}
func FindGroupbuy(c echo.Context) error {

	lists, err := models.AllOpenGroupbuy()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "groupbuy find issue")
	}

	return c.JSON(http.StatusOK, lists)
}
func AddIndividualOrder(c echo.Context) error {
	// get user
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	// get user
	user := claims["name"].(string)
	fmt.Print(user)
	// retrieve data posted

	var json AddIndividualOrderConfig
	if err := c.Bind(&json); err != nil {
		return err
	}

	order_id, err := models.CreateOrderByUser(json.Order[0].Groupbuy_id, user, json.Address)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "create order issues")
	}

	Itemerr := models.CreateUserOrderItems(user, order_id, json.Order)
	if Itemerr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "create groupbuy items issues")
	}

	return c.JSON(http.StatusOK, "order created!")
}

func ViewIndividualGroupbuy(c echo.Context) error {
	groupbuy_id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(1)
	fmt.Println(groupbuy_id)
	groupbuyInfo, err := models.GetSingleGroupbuy(groupbuy_id)
	// get groupbuy info(from there get user_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "view individual groupbuy issues")
	}
	user_id := groupbuyInfo.User_id

	itemInfo, itemErr := models.GetHostItems(user_id, groupbuy_id)
	if itemErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "view hostitem issues")
	}
	var individual ViewIndividualGroupbuyConfig
	individual.GroupbuyInfo = groupbuyInfo
	individual.HostItemInfo = itemInfo
	fmt.Println(user_id)
	return c.JSON(http.StatusOK, individual)
	// get items created by the creator according to group buy id

}
