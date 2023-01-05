/*
 * skogul, length sender
 *
 * Copyright (c) 2022 Telenor Norge AS
 * Author(s):
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package main

import (
	"fmt"

	"github.com/telenornms/skogul"
	"github.com/telenornms/skogul/sender"
)

func init() {
	sender.Auto.Add(skogul.Module{
		Name:    "length",
		Aliases: []string{"metriclength"},
		Alloc:   func() interface{} { return &Length{} },
		AutoMake: true,
		Help:    "Dummy sender that just prints the amount of metrics in a container. Tech/plugin demo.",
	})
}

/*
Length sender simply prints the amount of metrics per container
*/
type Length struct {}

// Send prints the JSON-formatted container to stdout
func (l *Length) Send(c *skogul.Container) error {
	fmt.Printf("length: %d\n", len(c.Metrics))
	return nil
}
